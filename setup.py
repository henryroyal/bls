from setuptools import setup, find_packages
from codecs import open
from os import path

here = path.abspath(path.dirname(__file__))

with open(path.join(here, 'README.md'), encoding='utf-8') as f:
    long_description = f.read()

with open(path.join(here, 'requirements.txt'), encoding='utf-8') as f:
    requirements = f.read()


setup(
    name='etlBLS',
    version='0.0.1',
    description='BLS ETL system',
    long_description=long_description,
    author='Henry Royal',
    author_email='dev@hwr.io',
    license='MIT',
    classifiers=[
        'Development Status :: 3 - Alpha',
        'License :: OSI Approved :: MIT License',
        'Programming Language :: Python :: 3.6',
    ],
    keywords='BLS ETL',
    packages=find_packages(),
    install_requires=requirements,
    entry_points={
        'console_scripts': ['extract_series_indices=bin.extract_series_indices:main'],
    }
)
