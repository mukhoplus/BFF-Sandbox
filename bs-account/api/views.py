import json
import os
from django.http import JsonResponse, HttpResponseNotFound

BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

def get_user_by_id(request, id):
    with open(os.path.join(BASE_DIR, 'db.json')) as f:
        data = json.load(f)
    
    user = next((user for user in data if user['id'] == id), None)
    
    if user is not None:
        return JsonResponse(user)
    else:
        return HttpResponseNotFound(f'User with id {id} not found')
