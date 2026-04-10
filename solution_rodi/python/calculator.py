# api.py
from flask import Flask, jsonify
import math

app = Flask(__name__)

def parse(x):
    return float(x)

@app.route('/sin/<x>')
def sin(x):
    x = parse(x)
    return jsonify(calculation=f"sin({x})", result=math.sin(x))

@app.route('/arcsin/<x>')
def arcsin(x):
    x = parse(x)
    if x < -1 or x > 1:
        return jsonify(error="arcsin domain is [-1,1]"), 400
    return jsonify(calculation=f"arcsin({x})", result=math.asin(x))

@app.route('/cos/<x>')
def cos(x):
    x = parse(x)
    return jsonify(calculation=f"cos({x})", result=math.cos(x))

@app.route('/arccos/<x>')
def arccos(x):
    x = parse(x)
    if x < -1 or x > 1:
        return jsonify(error="arccos domain is [-1,1]"), 400
    return jsonify(calculation=f"arccos({x})", result=math.acos(x))

@app.route('/tan/<x>')
def tan(x):
    x = parse(x)
    return jsonify(calculation=f"tan({x})", result=math.tan(x))

@app.route('/arctan/<x>')
def arctan(x):
    x = parse(x)
    return jsonify(calculation=f"arctan({x})", result=math.atan(x))

app.run(host='0.0.0.0', port=4000)
