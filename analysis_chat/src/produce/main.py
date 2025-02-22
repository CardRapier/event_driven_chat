import asyncio
import os

from ..pub_sub.nats_subscriber  import NatsSubscriber

def main():
    # Create and set the event loop
    loop = asyncio.new_event_loop()
    asyncio.set_event_loop(loop)

    nats_uri = os.getenv("NATS_URL", "localhost:4222")
    nats_pub = NatsSubscriber(loop, server_url=nats_uri)
    loop.run_until_complete(nats_pub.connect())