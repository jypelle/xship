# xship

A shoot'em up built with [tinygo](https://tinygo.org/) for the [pygamer](https://www.adafruit.com/product/4242)

![screenshoot1](images/screenshoot1.jpg)

*Maybe a good template to learn tinygo with microcontrollers*.

## Setup (Linux)

Install tinygo >= 0.16.0 (need this [fix](https://github.com/tinygo-org/tinygo/commit/db27541b1a44a903feeeef91840314a56fcdc725) to use TFT screen).

Connect your Pygamer to a USB port.

```bash
git clone git@github.com:jypelle/xship.git
cd xship
go get -u tinygo.org/x/drivers
make install
```

Enjoy!

