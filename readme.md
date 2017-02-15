# GoBot

My experimental Discord bot

### Cross Compiling
The gopus library uses the CGO package, therefore building both Mac and Linux require `CGO_ENABLED=1`.

> Work in progress

Other libraries are needed in order for CGO work properly cross platform.

```
sudo apt-get install gcc-multilib
apt-get install libpango1.0
```

This are some packages I came across as suggestions after a little browsing. Doesn't seem to be working currently.
This is something I need to investigate more on in the future.
