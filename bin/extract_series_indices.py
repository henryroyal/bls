import sys
import json
import click

import util.log
from bls.datasets import timeseries


@click.command()
@click.option('--outfile', '-o', type=click.File(mode='w'))
@click.option('--logging-level', envvar='LOGGING_LEVEL', type=util.log.LOG_LEVELS, default='WARNING')
@click.option('--logging-format', envvar='LOGGING_FORMAT', type=util.log.LOG_FORMATS, default='JSON')
def main(outfile):
    for series in timeseries:
        index = series.get_series_index()
        for item in index:
            item = {k: str(v).strip() for k, v in item.items()}
            outfile.write(json.dumps(item) + '\n')
            del item

        del series

    sys.exit(0)


if __name__ == '__main__':
    main()
