from django.shortcuts import render, redirect
import requests
import json 
import math
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
    "turn": 73,
    "x": 20,
    "y": 40
  },
  {
    "capacity": 20,
    "description": "water",
    "duration": 6,
    "name": "att2",
    "nextTurn": 501,
    "turn": 781,
    "x": 80,
    "y": 120
  },
  {
    "capacity": 1,
    "description": "Earth",
    "duration": 1,
    "name": "att3",
    "nextTurn": 10,
    "turn": 86,
    "x": 100,
    "y": 200
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
    res = requests.get('http://127.0.0.1:3000/API/rewardinfo/test')
    response = json.loads(res.text)
    print(response)
    return render(request, "test.html")

def atracciones(request):
    #global atracciones2
    res = requests.get('http://127.0.0.1:3000/attractions')
    response = json.loads(res.text)
    print(response)
    return render(request, "atracciones.html", {"atracciones":response})

def pqr(request):
    return render(request, "pqr.html")

def premios(request):
    #global premios2
    res = requests.get('http://127.0.0.1:3000/rewards')
    response = json.loads(res.text)
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
    result = 0
    for person in personas:
        print(person['id'])
        print(request.POST["id"])
        if int(person['id']) == int(request.POST["id"]):
            result = person
    res = requests.get('http://127.0.0.1:3000/api/userinfo/'+request.POST["id"])
    response = json.loads(res.text)
    print(response)
    
    if result != 0:
        return usuario_info(request, result)
    else:
        return usuario_info(request, 0)


def atraccion(request, nombre):
    # for a in atracciones2:
    #     if(a['name'] == nombre):
    #         atrac = a
    #         break
    #next_turn = atrac['nextTurn']+atrac['capacity']
    # print(atrac['turn']-atrac['nextTurn'])
    # time = (math.floor((atrac['turn']-atrac['nextTurn'])/atrac['capacity'])*atrac['duration'])
    res = requests.get('http://127.0.0.1:3000/api/attractioninfo/'+nombre)
    response = json.loads(res.text)
    next_turn = response['nextTurn']+response['capacity']
    time = (math.floor((response['currentTurn']-response['nextTurn'])/response['capacity'])*response['duration'])
    print(response)
    if time < 0:
        time = 0
    print(time)
    return render(request, "atraccion.html", {"atraccion":response, "next_turn":next_turn, "time":time})

def premio(request, nombre):
    # for a in premios2:
    #     if(a['name'] == nombre):
    #         prem = a
    #         break
    res = requests.get('http://127.0.0.1:3000/api/rewardinfo/'+nombre)
    response = json.loads(res.text)
    print(response)
    return render(request, "premio.html", {"premio":response})


def mapa(request):
    time = []
    res = requests.get('http://127.0.0.1:3000/attractions')
    response = json.loads(res.text)
    print(response)
    for atrac in response:
        time.append(math.floor((atrac['currentTurn']-atrac['nextTurn'])/atrac['capacity'])*atrac['duration'])
    return render(request, "mapa.html", {"atracciones": response, "times":time})


def pedir_turno(request, nombre):
    # result = 0
    # for atrac in atracciones2:
    #     if atrac['name'] == nombre:
    #         at = atrac
    #         break
    # for person in personas:
    #     print(person['id'])
    #     print(request.POST["id"])
    #     if int(person['id']) == int(request.POST["id"]):
    #         person['turn'] = atrac['turn']
    #         person['atraction'] = atrac['name']
    #         result = person
    params ={
        "Id":int(request.POST["id"]),
        "Attraction": nombre,
    }
    res = requests.post("http://127.0.0.1:3000/api/usernextturn/", json=params)
    response = json.loads(res.text)
    print(response)
    if response['message'] == 'Registro exitoso':
        res = requests.get('http://127.0.0.1:3000/api/userinfo/'+request.POST["id"])
        response = json.loads(res.text)
        return usuario_info(request, response)
    else:
        return usuario_info(request, 0)