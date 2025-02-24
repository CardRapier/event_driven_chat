import asyncio
import os

from ..pub_sub.channels import Channels
from ..pub_sub.nats_subscriber import NatsSubscriber
from ..pub_sub.nats_publisher import NatsPublisher
from .redpanda_consumer import RedpandaConsumer


async def main():
    nats_uri = os.getenv("NATS_URL", "localhost:4222")
    nats_sub = NatsSubscriber(server_url=nats_uri)
    await nats_sub.connect()

    # Instantiate and connect the NATS publisher
    nats_pub = NatsPublisher(server_url=nats_uri)
    await nats_pub.connect()
    # Instantiate the Redpanda consumer, injecting the NATS publisher
    consumer = RedpandaConsumer(nats_pub)
    producer = None

    def test(data):
        print(data)

    await nats_sub.subscribe(Channels.SENTIMENT_MESSAGE, test)

    try:
        await consumer.run()
    except KeyboardInterrupt:
        pass
    finally:
        await nats_pub.close()


if __name__ == "__main__":
    asyncio.run(main())
