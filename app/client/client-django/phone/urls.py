from django.urls import path
from phone import views
from django.conf import settings

urlpatterns = [
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