
import os
from dotenv import load_dotenv

load_dotenv()

class RedPandaConsumerConfig:
    redpanda_url: str

    def __init__(self, redpanda_url, group_id, subscribe=["chat"]):
        self.redpanda_url = redpanda_url
        self.group_id = group_id
        self.subscribe = subscribe

redpanda_consumer_config = RedPandaConsumerConfig(
    os.getenv("REDPANDA_URL", "localhost:9092"), "analysis-group"
)


class RedPandaProducerConfig:
    redpanda_url: str

    def __init__(self, redpanda_url, topic="chat"):
        self.redpanda_url = redpanda_url
        self.topic = topic


redpanda_producer_config = RedPandaProducerConfig(
    os.getenv("REDPANDA_URL", "localhost:9092")
)