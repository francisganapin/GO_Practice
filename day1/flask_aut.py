from flask import Flask,jsonify,request
from flask_jwt_extended import create_access_token, jwt_required, get_jwt_identity, JWTManager



app = Flask(__name__)
app.config["JWT_SECRET_KEY"] = "supersecretkey"
jwt = JWTManager(app)

users = {"admin":"password123"}

@app.route('/login',methods=['POST'])
def login():
    data = request.json
    username = data.get('username')
    password = data.get('password')

    if username in users and users[username] == password:
        access_token = create_access_token(identity=username)
        return jsonify(access_token=access_token)
    return jsonify({'message':'Invalid credential'}),401


@app.route("/proctected",methods=["GET"])
@jwt_required()
def protected():
    current_user = get_jwt_identity()
    return jsonify(logged_in_as=current_user)


if __name__ == '__main__':
    app.run(debug=True)