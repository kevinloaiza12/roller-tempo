from django.shortcuts import render, redirect
import math
import requests, json

from django.conf import settings
from django.templatetags.static import static
# Create your views here.

back_server="http://35.184.16.30:3000"

personas = [
    {
     "coins": 50,
     "id": 1012002,
     "turn": 0,
     "atraction": "NA"
    },
    {
     "coins": 60,
     "id": 123456789,
     "turn": 0,
     "atraction": "NA"
    },
    {
     "coins": 20,
     "id": 987654321,
     "turn": 0,
     "atraction": "NA"
    }
]


atracciones2 = [
  {
    "capacity": 50,
    "description": "Fire",
    "duration": 60,
    "name": "att1",
    "nextTurn": 50,
    "currentTurn": 73,
    "x": 20,
    "y": 40,
    "image": "logo.png"
  },
  {
    "capacity": 20,
    "description": "water",
    "duration": 6,
    "name": "att2",
    "nextTurn": 501,
    "currentTurn": 781,
    "x": 80,
    "y": 120,
    "image": "placeholder.jpg"
  },
  {
    "capacity": 1,
    "description": "Earth",
    "duration": 1,
    "name": "att3",
    "nextTurn": 10,
    "currentTurn": 86,
    "x": 100,
    "y": 200,
    "image": "premio_placeholder.jpg"
  }
]

def atraccion(request, nombre):
    res = requests.get('http://35.184.16.30:3000/api/attractions/'+nombre)
    response = json.loads(res.text)
    next_turn = response['NextTurn']+response['Capacity']
    time = (math.floor((response['CurrentRoundTurn']-response['NextTurn'])/response['Capacity'])*response['Duration'])
    if time < 0:
        time = 0
    print(time)
    last_turn = response['CurrentRoundTurn']+response['Capacity']
    return render(request, "Attr_default.html", {"atraccion":response, "last_turn":last_turn, "next_turn":next_turn, "time":time, "id":nombre})


def registrar_turno(request, nombre):
    res = requests.get('http://35.184.16.30:3000/api/attractions/'+nombre)
    response = json.loads(res.text)
    next_turn = response['NextTurn']+response['Capacity']
    time = (math.floor((response['CurrentRoundTurn']-response['NextTurn'])/response['Capacity'])*response['Duration'])
    if time < 0:
        time = 0
    print(time)
    turns = [*range(response['CurrentRoundTurn'], response['CurrentRoundTurn']+response['Capacity'])]
    return render(request, "Attr_register.html", {"atraccion":response, "turns":turns, "next_turn":next_turn, "time":time, "id":nombre})


def usar_turno(request, nombre):
    res = requests.get('http://35.184.16.30:3000/api/attractions/'+nombre)
    response = json.loads(res.text)
    next_turn = response['NextTurn']+response['Capacity']
    time = (math.floor((response['CurrentRoundTurn']-response['NextTurn'])/response['Capacity'])*response['Duration'])
    if time < 0:
        time = 0
    print(time)
    turns = [*range(response['CurrentRoundTurn'], response['CurrentRoundTurn']+response['Capacity'])]
    return render(request, "Attr_turn.html", {"atraccion":response, "turns":turns, "next_turn":next_turn, "time":time, "id":nombre})


def registrar(request):
    response = 0
    print((request.POST["attr"]))
    print((request.POST["id"]))
    params ={
        "UserID":int(request.POST["id"]),
        "AttractionID": int(request.POST["attr"]),
    }
    res = requests.post("http://35.184.16.30:3000/api/users/turn", json=params)
    response = json.loads(res.text)
    print(response)
    print("hola")
    if'error' in response:
        return redirect('register', nombre=request.POST['attr'])
    else:
        return redirect('attraction', nombre=request.POST['attr'])
        
    

def usar_t(request):
    res_user = requests.get('http://35.184.16.30:3000/api/users/'+request.POST["id"])
    res_user = json.loads(res_user.text)
    res_attraction = requests.get('http://35.184.16.30:3000/api/attractions/'+request.POST["attr"])
    res_attraction = json.loads(res_attraction.text)
    print(res_attraction['CurrentRoundTurn'])
    print(res_attraction['CurrentRoundTurn']+res_attraction['Capacity'])
    if(int(res_user['Attraction']) == int(request.POST["attr"])and (res_user['Turn'] >= res_attraction['CurrentRoundTurn'] and res_user['Turn'] <= res_attraction['CurrentRoundTurn'] +res_attraction['Capacity'])):
        print("paaas")
        requests.put("http://35.184.16.30:3000/api/users/"+ request.POST["id"] +"/removeturn")
        requests.post("http://35.184.16.30:3000/api/users/"+ request.POST["id"] +"/reward?amount=5")
        return redirect('attraction', nombre=request.POST['attr'])

    return redirect('use', nombre=request.POST['attr'])

        
def rewards(request):
    global premios2
    res = requests.get('http://35.184.16.30:3000/api/rewards')
    response = json.loads(res.text)['message']
    print(response)
    return render(request, "rewards.html", {"premios":response})


def reward(request, nombre):
    res = requests.get('http://35.184.16.30:3000/api/rewards/'+nombre)
    response = json.loads(res.text)
    print(response)
    return render(request, "reward.html", {"premio":response, "id":nombre})


def buy_reward(request):
    params ={
        "UserID":int(request.POST["user_id"]),
        "RewardID": int(request.POST["reward_id"]),
    }
    res = requests.post("http://35.184.16.30:3000/api/users/buyreward", json=params)
    response = json.loads(res.text)
    print(response)
    print("hola")
    if'error' in response:
        return redirect('reward', nombre=request.POST['reward_id'])
    else:
        return redirect('rewards')