"""rollertempo URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/4.1/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path
from phone import views

urlpatterns = [
    path('admin/', admin.site.urls),
    path('phone/', views.phone, name='phone'),
    path('pqr/', views.pqr, name='pqr'),
    path('atracciones/', views.atracciones, name='atracciones'),
    path('premios/', views.premios, name='premios'),
    path('usuario/', views.usuario, name='usuarios'),
    path('usuario/<slug:nombre>', views.usuario, name='usuarios'),
    path('atraccion/<slug:nombre>', views.atraccion, name='atraccion'),
    path('premio/<slug:nombre>', views.premio, name='premio'),
    path('buscar/', views.buscar, name='buscar'),
    path('mapa/', views.mapa, name='mapa'),
    path('pedir_turno/<slug:nombre>', views.pedir_turno, name='pedir_turno'),
]
