import asyncio
import os

from src.pub_sub.nats_publisher import NatsPublisher

from ..analysis.main import Analysis
from .manager import Manager

from ..pub_sub.channels import Channels
from ..pub_sub.nats_subscriber import NatsSubscriber

from dotenv import load_dotenv

load_dotenv()


def main():
    # Create and set the event loop
    loop = asyncio.new_event_loop()
    asyncio.set_event_loop(loop)

    nats_uri = os.getenv("NATS_URL", "localhost:4222")
    nats_sub = NatsSubscriber(loop, server_url=nats_uri)
    loop.run_until_complete(nats_sub.connect())

    analysis = Analysis()
    nats_pub = NatsPublisher(loop, server_url=nats_uri)
    loop.run_until_complete(nats_pub.connect())

    manager = Manager(publisher=nats_pub, analysis=analysis)

    loop.run_until_complete(
        nats_sub.subscribe(Channels.CHAT_MESSAGE, manager.manage_message)
    )

    # Keep the event loop running to listen for incoming messages
    try:
        loop.run_forever()
    except KeyboardInterrupt:
        print("Shutting down...")
    finally:
        loop.run_until_complete(nats_sub.close())
        loop.close()


if __name__ == "__main__":
    main()
