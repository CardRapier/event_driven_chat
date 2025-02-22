
import os
from dotenv import load_dotenv

load_dotenv()

print(f"Redpanda URI: {os.getenv("REDPANDA_URL")}")

class RedPandaConfig():
    redpanda_url: str

    def __init__(self, redpanda_url, group_id, subscribe=["chat"]):
        self.redpanda_url = redpanda_url
        self.group_id = group_id
        self.subscribe = subscribe

redpanda_consumer_config = RedPandaConfig(
    os.getenv("REDPANDA_URL", "localhost:9092"),
    "analysis-group"
)