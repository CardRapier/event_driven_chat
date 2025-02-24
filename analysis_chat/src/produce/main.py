import asyncio
import os

from src.pub_sub.nats_publisher import NatsPublisher

from ..analysis.main import Analysis
from .manager import Manager

from ..pub_sub.channels import Channels
from ..pub_sub.nats_subscriber import NatsSubscriber

from dotenv import load_dotenv

load_dotenv()


async def main():
    # Create and set the event loop

    nats_uri = os.getenv("NATS_URL", "localhost:4222")
    nats_sub = NatsSubscriber(server_url=nats_uri)
    await nats_sub.connect()

    analysis = Analysis()
    nats_pub = NatsPublisher(server_url=nats_uri)
    await nats_pub.connect()

    manager = Manager(publisher=nats_pub, analysis=analysis)

    await nats_sub.subscribe(Channels.CHAT_MESSAGE, manager.manage_message)

    # Keep the event loop running to listen for incoming messages
    stop_event = asyncio.Event()
    try:
        await stop_event.wait()
    except KeyboardInterrupt:
        print("Shutting down...")
    finally:
        await nats_sub.close()
        await nats_pub.close()


if __name__ == "__main__":
    asyncio.run(main())