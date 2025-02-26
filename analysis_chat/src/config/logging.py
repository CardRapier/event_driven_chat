import logging


class CustomFormatter(logging.Formatter):
    """
    Custom logging formatter to add colors and limit the logger name
    to the final two components.
    """

    # ANSI escape codes for colors
    COLORS = {
        "DEBUG": "\033[37m",  # White
        "INFO": "\033[32m",  # Green
        "WARNING": "\033[33m",  # Yellow
        "ERROR": "\033[31m",  # Red
        "CRITICAL": "\033[41m",  # Red background
    }
    RESET = "\033[0m"

    def format(self, record: logging.LogRecord) -> str:

        # Determine the color for this log level
        levelname = logging.getLevelName(record.levelno)
        color = self.COLORS.get(levelname, "")
        warning = self.COLORS.get("WARNING", "")

        # Color the levelname with its corresponding color and wrap it in brackets
        record.levelname = f"{color}{levelname}{self.RESET}"

        # Color the message. Note: Use record.getMessage() to get the unformatted message.
        record.msg = f"{color}{record.getMessage()}{self.RESET}"

        # Color the logger name
        record.name = f"{warning}[{record.module}]{self.RESET}"

        # Now let the base class format the record
        return super().format(record)

    def formatTime(self, record: logging.LogRecord, datefmt=None) -> str:
        # Use the base class to generate the timestamp string
        return super().formatTime(record, datefmt)


# Set up the logger
logger = logging.getLogger("analysis_chat")
logger.setLevel(logging.DEBUG)

# Create a stream handler for console output
stream_handler = logging.StreamHandler()
stream_handler.setLevel(logging.DEBUG)

# Define a format string. This will include the time, colored level name,
# trimmed and colored logger name, and the colored message.
fmt = "%(asctime)s %(levelname)s %(name)s %(message)s"
datefmt = "%d/%m/%Y, %H:%M:%S"
formatter = CustomFormatter(fmt, datefmt)
stream_handler.setFormatter(formatter)

# Attach the handler to the logger
logger.addHandler(stream_handler)
