import asyncio
from analysis_chat.src.config.logging import logger
from typing import Any, Optional
from confluent_kafka import Consumer  # type: ignore
from analysis_chat.src.events import Events
from analysis_chat.src.pub_sub.nats_publisher import NatsPublisher
from analysis_chat.src.config.redpanda import redpanda_consumer_config
from analysis_chat.src.pub_sub.channels import Channels

class RedpandaConsumer:
    def __init__(self, nats_publisher: NatsPublisher):
        self.nats_publisher = nats_publisher
        config = {
            "bootstrap.servers": redpanda_consumer_config.redpanda_url,
            "group.id": redpanda_consumer_config.group_id,
            "auto.offset.reset": "earliest",
        }
        self.consumer = Consumer(config)
        # redpanda_consumer_config.subscribe should be a list of topics
        self.consumer.subscribe(redpanda_consumer_config.subscribe)

    async def run(self) -> None:
        try:
            while True:
                msg: Optional[Any] = await asyncio.to_thread(self.consumer.poll, 1.0)
                if msg is None:
                    continue
                if msg.error():
                    logger.error(f"Consumer error: {msg.error()}")
                    continue

                # Check if the message key exists and decode it.
                key_bytes: Optional[bytes] = msg.key()
                key: str = key_bytes.decode("utf-8") if key_bytes is not None else ""
                if key != Events.CHAT_MESSAGE:
                    continue

                message_data: str = msg.value().decode("utf-8")
                logger.info(f"Received message from Redpanda: {message_data}")

                try:
                    # Publish to NATS using the injected publisher
                    await self.nats_publisher.publish(
                        Channels.CHAT_MESSAGE, message_data
                    )
                except Exception as e:
                    logger.error("Error publishing to NATS:", e)
        except KeyboardInterrupt:
            logger.warning("Shutting down consumer...")
        finally:
            self.consumer.close()