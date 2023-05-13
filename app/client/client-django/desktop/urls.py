from django.urls import path
from desktop import views
from django.conf import settings

urlpatterns = [
    path('desktop/attractiones/registrar/', views.registrar, name="register2"),
    path('desktop/attractiones/usar/', views.usar_t, name="use2"),
    path('desktop/attractiones/register/<slug:nombre>/', views.registrar_turno, name="register"),
    path('desktop/attractiones/use/<slug:nombre>/', views.usar_turno, name="use"),
    path('desktop/attraction/<slug:nombre>/', views.atraccion, name="attraction"),
]