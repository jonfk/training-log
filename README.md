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

##List of workout names

###Squats:
- high bar squats
- low bar squats
- front squats
- paused high bar squats
- paused low bar squats
- paused front squats
- db lunges

###Pressing:
- close grip bench press
- bench press
- tng bench press
- overhead press
- behind the neck press
- db incline press
- db flyes

###Pulling:
- sumo deadlift
- conventional deadlift
- deficit conventional deadlift
- stiff leg deadlift
- block pulls
- sumo block pulls
- bent over rows
- pendlay rows
- chest supported rows

###Back:
- pull ups
- chin ups
- lat pulldowns

###Arms:
- alternating db curls