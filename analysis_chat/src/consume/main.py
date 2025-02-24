import asyncio
import os
import threading
from ..pub_sub.nats_publisher import NatsPublisher
from .redpanda_consumer import RedpandaConsumer


def main():
    # Create and set the event loop
    loop = asyncio.new_event_loop()
    asyncio.set_event_loop(loop)

    # Instantiate and connect the NATS publisher
    nats_uri = os.getenv("NATS_URL", "localhost:4222")
    nats_pub = NatsPublisher(loop, server_url=nats_uri)
    loop.run_until_complete(nats_pub.connect())

    loop_thread = threading.Thread(target=loop.run_forever, daemon=True)
    loop_thread.start()
    # Instantiate the Redpanda consumer, injecting the NATS publisher
    consumer = RedpandaConsumer(nats_pub)

    try:
        consumer.run()
    except KeyboardInterrupt:
        pass
    finally:
        loop.call_soon_threadsafe(loop.stop)
        loop_thread.join()
        loop.run_until_complete(nats_pub.close())
        loop.close()


if __name__ == "__main__":
    main()
