training-log
============

my training log in yaml format

##Template
```yaml
date: yyyy-mm-dd
time: hh:mmPM                                   #OPTIONAL key
length: 3h                                      #OPTIONAL key
bodyweight: 68.0kg                              #OPTIONAL key
event: Saguenay Open et Provincial Equipe 2014  #OPTIONAL key   Event key
wilks: 290.41                                   #OPTIONAL key   Event key
total: 360.0 kg                                 #OPTIONAL key   Event key
workout:
  - name: squats
    weight: 225 lbs                             #OPTIONAL key
    sets: 5
    reps: 5
    exertion: rpe 8                             #OPTIONAL key
notes:
  - ...
```