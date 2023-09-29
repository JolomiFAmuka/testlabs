
from . import views
from django.urls import path

urlpatterns = [
    path("", views.index, name="index"),
    path("item/", views.item, name="item"),
    path("uwale/", views.uwale, name="uwale"),
    path("uranran", views.uranran, name="uranran")
]
