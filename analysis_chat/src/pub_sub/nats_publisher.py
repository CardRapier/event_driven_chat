from analysis_chat.src.config.logging import logger
from nats.aio.client import Client as NATS
class NatsPublisher:
    def __init__(self, server_url="nats://localhost:4222"):
        self.server_url = server_url
        self.nc = NATS()

    async def connect(self):
        await self.nc.connect(self.server_url)
        logger.info("Connected to NATS")

    async def publish(self, subject, message):
        await self.nc.publish(subject, message.encode("utf-8"))
        await self.nc.flush()
        logger.info(f"Published message to subject '{subject}': {message}")

    async def close(self):
        await self.nc.close()
