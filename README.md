# Narr
Netflix Audio Ripper - Download audio tracks from Netflix to sample your favourite shows. :musical_note:

## Usage

Start Chrome or Brave with remote debugging enabled:

```bash
google-chrome --remote-debugging-port=9222
```

- Make sure that you are logged in on Netflix in the started browser.
- Navigate to the show you want to download the audio from.
- Copy the URL of the show.
- Run narr with the URL as argument.

```bash
narr "https://www.netflix.com/watch/12345678?trackId=12345678"
```

Observe the progress in the terminal:

```bash
2023/02/18 18:34:25 ▼ Downloading https://www.netflix.com/watch/81237996?trackId=14170056  ⟾  /home/looper/81237996-14170056-4037200794235010051
2023/02/18 18:34:35 ✓ Finished    https://www.netflix.com/watch/81237996?trackId=14170056  ⟾  /home/looper/81237996-14170056-4037200794235010051, got 65346400 bytes
```

You can navigate to any other show or episode or change the language of the audio track while narr is running. It will
download the audio track of the currently playing episode.

```bash
2023/02/18 18:34:25 ▼ Downloading https://www.netflix.com/watch/81237996?trackId=14170056  ⟾  /home/looper/81237996-14170056-4037200794235010051
2023/02/18 18:34:35 ✓ Finished    https://www.netflix.com/watch/81237996?trackId=14170056  ⟾  /home/looper/81237996-14170056-4037200794235010051, got 65346400 bytes
2023/02/18 16:59:42 🗺Navigate to https://www.netflix.com/watch/81238005?trackId=14170056 
2023/02/18 16:59:43 ▼ Downloading https://www.netflix.com/watch/81238005?trackId=14170056  ⟾  /home/looper/81238005-14170056-605394647632969758
2023/02/18 16:59:53 ✓ Finished    https://www.netflix.com/watch/81238005?trackId=14170056  ⟾  /home/looper/81238005-14170056-605394647632969758, got 65346400 bytes
```

It is also possible to navigate to the netflix home page. Narr will then download audio tracks of trailers or previews.

# How it works

Narr uses the [Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/) to communicate with the
browser. It intercepts the network requests
and downloads the audio tracks in mp4 format. This is possible because the audio is not protected by DRM.

# Build from source

Narr is written in Go. You can build it the following commands:

```bash
git clone github.com/IljaN/narr
cd narr
go build
```

# Flags

```bash
narr --help
~/c/narr:λ ./narr --help
Usage: narr [--chrome-url CHROME-URL] VIDEOURL [DOWNLOADDIR]

Positional arguments:
  VIDEOURL               url of the video to download audio from. Must be a netflix url. e.g. https://www.netflix.com/watch/12345678?trackId=12345678
  DOWNLOADDIR            directory where to download the audio files. Defaults to current working directory.

Options:
  --chrome-url CHROME-URL, -c CHROME-URL
                         url of the chrome debugger. [default: http://127.0.0.1:9222]
  --help, -h             display this help and exit
``
