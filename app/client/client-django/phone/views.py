from django.shortcuts import render, redirect
import requests
import json 
import math
personas = [
    {
     "coins": 50,
     "id": 1012002,
     "turn": 0,
     "attraction": "NA"
    },
    {
     "coins": 60,
     "id": 123456789,
     "turn": 0,
     "attraction": "NA"
    },
    {
     "coins": 20,
     "id": 987654321,
     "turn": 0,
     "attraction": "NA"
    }
]

atracciones2 = [
    {
    "capacity": 50,
    "description": "Fire",
    "duration": 60,
    "name": "5",
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
premios2 = [
  {
    "price": "50",
    "description": "Fire",
    "name": "atraccion",
  },
  {
    "description": "water",
    "price": "6",
    "name": "atracci",

  },
  {
    "description": "Earth",
    "price": "1",
    "name": "atrac",
  }
]

# Create your views here.
def phone(request):
    return render(request, "test.html")

def atracciones(request):
    global atracciones2
    res = requests.get('http://127.0.0.1:3000/api/attractions')
    response = json.loads(res.text)
    print(response['message'])
    return render(request, "atracciones.html", {"atracciones":response['message']})

def pqr(request):
    return render(request, "pqr.html")

def premios(request):
    global premios2
    res = requests.get('http://127.0.0.1:3000/api/rewards')
    response = json.loads(res.text)['message']
    print(response)
    return render(request, "premios.html", {"premios":response})

def usuario(request):
    return render(request, "usuario.html")


def usuario_info(request, data):
    print(data)
    if data != 0:
      return render(request, "usuario.html", {"data":data})
    else:
      return render(request, "usuario.html")
    
def buscar(request):
    response = 0
    res = requests.get('http://127.0.0.1:3000/api/users/'+request.POST["id"])
    response = json.loads(res.text)
    print(response)
    
    if response != 0:
        return usuario_info(request, response)
    else:
        return usuario_info(request, 0)


def atraccion(request, nombre):
    res = requests.get('http://127.0.0.1:3000/api/attractions/'+nombre)
    response = json.loads(res.text)
    next_turn = response['NextTurn']+response['Capacity']
    time = (math.floor((response['CurrentRoundTurn']-response['NextTurn'])/response['Capacity'])*response['Duration'])
    if time < 0:
        time = 0
    print(time)
    return render(request, "atraccion.html", {"atraccion":response, "next_turn":next_turn, "time":time, "id":nombre})

def premio(request, nombre):
    res = requests.get('http://127.0.0.1:3000/api/rewards/'+nombre)
    response = json.loads(res.text)
    print(response)
    return render(request, "premio.html", {"premio":response})


def mapa(request):
    time = []
    response2 = []
    res = requests.get('http://127.0.0.1:3000/api/attractions')
    response = json.loads(res.text)['message']
    print(response)
    for atrac in response:
        t = (math.floor((atrac['CurrentRoundTurn']-atrac['NextTurn'])/atrac['Capacity'])*atrac['Duration'])
        if t > 0:
            time.append(t)
        else:
            time.append(0)
        del atrac["DeletedAt"]
    return render(request, "mapa.html", {"atracciones": response, "times":time})


def pedir_turno(request, nombre):
    response = 0
    print(nombre)
    print(int(request.POST["id"]))
    params ={
        "UserID":int(request.POST["id"]),
        "AttractionID": int(nombre),
    }
    res = requests.post("http://127.0.0.1:3000/api/users/turn", json=params)
    response = json.loads(res.text)
    print(response)
    print("hola")
    if'error' in response:
        return redirect('atraccion', nombre=nombre)
    else:
        res = requests.get('http://127.0.0.1:3000/api/users/'+request.POST["id"])
        response = json.loads(res.text)
        return usuario_info(request, response)
        