# Chrome

Chrome not loading up because it thinks it is running:

```shell
sudo rm -f ~/.config/google-chrome/SingletonLock
sudo rm -f ~/.config/google-chrome/SingletonCookie
sudo rm -f ~/.config/google-chrome/SingletonSocket
```