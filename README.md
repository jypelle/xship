# xship

A shoot'em up built with [tinygo](https://tinygo.org/) for the [pygamer](https://www.adafruit.com/product/4242)

![screenshoot1](images/screenshoot1.jpg)

*Maybe a good template to learn tinygo with microcontrollers*.

## Installation (Linux)

Install tinygo >= 0.16.0 (need this [fix](https://github.com/tinygo-org/tinygo/commit/db27541b1a44a903feeeef91840314a56fcdc725) to use TFT screen on pygamer).

Connect your pygamer to a USB port.

```bash
git clone git@github.com:jypelle/xship.git
cd xship
go get -u tinygo.org/x/drivers
make install
```

Enjoy!

### Installation with new game asset

If you have updated [game asset](images/asset.png):

Install go >= 1.9

```bash
make asset
make install
```
