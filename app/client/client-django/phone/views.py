from django.shortcuts import render, redirect
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

    return render(request, "test.html")

def atracciones(request):
    global atracciones2;
    return render(request, "atracciones.html", {"atracciones":atracciones2})

def pqr(request):
    return render(request, "pqr.html")

def premios(request):
    global premios2
    return render(request, "premios.html", {"premios":premios2})

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
    if result != 0:
        return usuario_info(request, result)
    else:
        return usuario_info(request, 0)


def atraccion(request, nombre):
    for a in atracciones2:
        if(a['name'] == nombre):
            atrac = a
            break
    next_turn = atrac['nextTurn']+atrac['capacity']
    print(atrac['turn']-atrac['nextTurn'])
    time = (math.floor((atrac['turn']-atrac['nextTurn'])/atrac['capacity'])*atrac['duration'])
    if time < 0:
        time = 0
    print(time)
    
    return render(request, "atraccion.html", {"atraccion":atrac, "next_turn":next_turn, "time":time})

def premio(request, nombre):
    for a in premios2:
        if(a['name'] == nombre):
            prem = a
            break
    return render(request, "premio.html", {"premio":prem})


def mapa(request):
    time = []
    for atrac in atracciones2:
        time.append(math.floor((atrac['turn']-atrac['nextTurn'])/atrac['capacity'])*atrac['duration'])
    return render(request, "mapa.html", {"atracciones": atracciones2, "times":time})


def pedir_turno(request, nombre):
    result = 0
    for atrac in atracciones2:
        if atrac['name'] == nombre:
            at = atrac
            break
    for person in personas:
        print(person['id'])
        print(request.POST["id"])
        if int(person['id']) == int(request.POST["id"]):
            person['turn'] = atrac['turn']
            person['atraction'] = atrac['name']
            result = person
    if result != 0:
        print(result)
        return usuario_info(request, result)
    else:
        return usuario_info(request, 0)
    