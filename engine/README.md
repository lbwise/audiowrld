FLOW:
![Screenshot 2025-11-15 at 10.36.21 am.png](../../../../../var/folders/85/67lzwrcx41g78m72j6hwrh9w0000gn/T/TemporaryItems/NSIRD_screencaptureui_7WzXYa/Screenshot%202025-11-15%20at%2010.36.21%E2%80%AFam.png)

Notes - future proofing:
- VSTs
- Stereo
- Piano roll
- One shot samples
- Loop samples
- Return tracks
- track routing
- automation

So pretty much the tick sequence should be like this:
On start we start the clock and put the midi scanner in the background
Then on tick:
1. Engine tick update data
2. Midi channel generate sound chunk
3. Audio tracks take sound from record and Midi tracks take audio from midi channel
4. Processing for each track
5. mixing
6. play and save audio