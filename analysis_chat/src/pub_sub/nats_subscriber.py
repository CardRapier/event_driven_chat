from nats.aio.client import Client as NATS

from analysis_chat.src.config.logging import logger


class NatsSubscriber:
    def __init__(self, server_url="nats://localhost:4222"):
        self.server_url = server_url
        self.nc = NATS()
        self.cbs = {}

    async def connect(self):
        await self.nc.connect(self.server_url)
        logger.info("Connected to NATS")

    async def subscribe(self, subject, cb):
        self.cbs[subject] = cb
        await self.nc.subscribe(subject, cb=self.message_handler)
        logger.info(f"Subscribe to subject '{subject}'")

    async def message_handler(self, msg):
        subject = msg.subject
        cb = self.cbs[subject]
        data = msg.data.decode("utf-8")
        logger.info(f"Received a message on '{subject}': {data}")
        await cb(msg)

    async def close(self):
        await self.nc.close()
