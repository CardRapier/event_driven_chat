from nats.aio.client import Client as NATS


class NatsSubscriber:
    def __init__(self, loop, server_url="nats://localhost:4222"):
        self.loop = loop
        self.server_url = server_url
        self.nc = NATS()

    async def connect(self):
        await self.nc.connect(self.server_url)
        print("Connected to NATS")

    async def subscribe(self, subject):
        await self.nc.subscribe(subject, self.message_handler)
        print(f"Subscribe to subject '{subject}'")

    async def message_handler(msg, cb):
        subject=msg.subject
        data = msg.data.decode("utf-8")
        print(f"Received a message on '{subject}': {data}")

    async def close(self):
        await self.nc.close()
