import json
from confluent_kafka import Producer  # type: ignore
from src.events import Events
from src.config.redpanda import redpanda_producer_config


class RedpandaProducer:
    def __init__(self):
        config = {
            "bootstrap.servers": redpanda_producer_config.redpanda_url,
        }
        self.producer = Producer(config)
        self.topic = redpanda_producer_config.topic

    async def publish(self, message):
        data = message.data.decode("utf-8")
        payload = json.dumps(data)
        print(f"Publishing message to Redpanda: {data}")
        try:
            self.producer.produce(self.topic, key=Events.SENTIMENT_MESSAGE, value=payload, callback=self.delivery_report)
            # Flush to ensure all messages are sent
            self.producer.flush()
        except Exception as e:
            print(f"Exception while producing message: {e}")

    def delivery_report(self, err, msg):
        """Called once for each message produced to indicate delivery result."""
        if err is not None:
            print(f"Message delivery failed: {err}")
        else:
            print(f"Message delivered to {msg.topic()} [{msg.partition()}] at offset {msg.offset()}")