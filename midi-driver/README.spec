Midi has three values as part of a message:

Note number:
0 - 127
C4 is 60
C0 - C7

Velocity:
0 - 127
Loudness of the note

Channel:
usually 1 - 16
channels should be memorized

Include NoteOn message: velocity > 0
and NoteOff message: velocity = 0
Note: One shot samples are triggered by NoteOn message only

Continuous controller messages (CC) - eg. sliders, benders:
0 - 127
- aftertouch (the pressure increasing after a note is hit
- pitch bend
