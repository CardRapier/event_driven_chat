import asyncio
from confluent_kafka import Consumer  # type: ignore
from src.config.redpanda import redpanda_consumer_config
from src.pub_sub.channels import Channels


class RedpandaConsumer:
    def __init__(self, nats_publisher: object):
        self.nats_publisher = nats_publisher
        config = {
            "bootstrap.servers": redpanda_consumer_config.redpanda_url,
            "group.id": redpanda_consumer_config.group_id,
            "auto.offset.reset": "earliest",
        }
        self.consumer = Consumer(config)
        # redpanda_consumer_config.subscribe should be a list of topics
        self.consumer.subscribe(redpanda_consumer_config.subscribe)

    def run(self):
        try:
            while True:
                msg = self.consumer.poll(timeout=1.0)
                if msg is None:
                    continue
                if msg.error():
                    print(f"Consumer error: {msg.error()}")
                    continue

                message_data = msg.value().decode("utf-8")
                print(f"Received message from Redpanda: {message_data}")

                # Publish to NATS using the injected publisher
                future = asyncio.run_coroutine_threadsafe(
                    self.nats_publisher.publish(Channels.CHAT_MESSAGE, message_data),
                    self.nats_publisher.loop,
                )
                try:
                    # Wait for the publishing coroutine to finish
                    future.result(timeout=5)
                except Exception as e:
                    print("Error publishing to NATS:", e)
        except KeyboardInterrupt:
            print("Shutting down consumer...")
        finally:
            self.consumer.close()
