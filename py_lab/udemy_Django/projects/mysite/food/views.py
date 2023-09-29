from django.shortcuts import render
from django.http import HttpResponse

# Create your views here.
def index(request):
    return HttpResponse("Hello Django Study")

def item(request):
    return HttpResponse("This is an item view")

def uwale(request):
    return HttpResponse("<h1>Uwale Special page</h1>")

def uranran(request):
    return HttpResponse("<h1>uranran Special page</h1>")