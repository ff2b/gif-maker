# GIF maker Help
This app is .gif file creation tool. Powered by Go lang UI framework fyne  

## How to use?
1. Open workspace folder. To open folder, You can click top left folder icon.
2. Choose folder, then click Open button.
3. .jpg and .png file list appeared, Check files which you want to select make gif frame images.
4. Click **Create GIF** button, and click **OK** button, then .gif file createtion start!
5. Complete creation gif, preview will display
6. You can save **Save GIF** button to save gif, or **Back Home** to back home view

## View Description
### 1.Home
Home view is displayed when you open this app.  

#### tooltip icons  
- openfolder: Open folder view open. You can choose workspace folder.  
- help: Help window open.  
- refresh: Refresh home view.

### 2. Preview
Created .gif file preview.  You can save or back home view.

#### tooltip icons
- play: Play or Replay GIF
- stop: Stop GIF playing

### 3. Options
Application option setting window. Click top menu **Edit - Options**, open option window.
If save setting, click **Apply** button.

#### options
- **GIF Settings**  
  - GIF frame rate: This value smaller, that frames switching faster.  
  - GIF loop: If check on, GIF play infinite loop.  
- **Application Settings**  
  - Default workspace folder: The workspace will open automatically from the next time.  
  - Default save folder: The save folder will set automatically, when you click Save GIF button at preview.  
  
NOTE: this options load from config/config.json  

## How to run
### Run from source code
Open termnal like Powershell/bash/mac terminal.  
Then execute that command.  
```
go run .
```

### Open from Compiled binary
Download Releases from this github project. Then unpack zip.  
Doubleclick **gif-maker** file  


## Disclaimer
- This App is experimental demo application.
- Generated GIF The file is probably of dirty quality, This is known issue.
- We are not responsible for any problems that occur using this app.
- Please raise an GitHub issue for bugs and improvement requests.
- I apologize for my poor English, but thank you so much for finding this App!.