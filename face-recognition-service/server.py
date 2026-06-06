from flask import Flask, request, jsonify
import insightface
import cv2
import numpy as np
import os
import json
from insightface.app import FaceAnalysis

app = Flask(__name__)

try:
    ctx = FaceAnalysis(allowed_modules=['detection', 'recognition'], providers=['CPUProvider'])
    ctx.prepare(ctx_id=0, det_size=(640, 480))
except Exception as e:
    print(f"Error initializing FaceAnalysis: {e}")
    ctx = None

@app.route('/api/embedding', methods=['POST'])
def generate_embedding():
    """Generate face embedding from image"""
    try:
        if 'image' not in request.files:
            return jsonify({'success': False, 'error': 'No image provided'}), 400

        file = request.files['image']
        
        img_data = file.read()
        img_array = np.frombuffer(img_data, np.uint8)
        img = cv2.imdecode(img_array, cv2.IMREAD_COLOR)

        if img is None:
            return jsonify({'success': False, 'error': 'Invalid image'}), 400

        faces = ctx.get(img)

        if not faces:
            return jsonify({'success': False, 'error': 'No face detected'}), 400

        embedding = faces[0]['embedding'].tolist()

        return jsonify({
            'success': True,
            'embedding': embedding
        })

    except Exception as e:
        return jsonify({'success': False, 'error': str(e)}), 500


@app.route('/api/detect', methods=['POST'])
def detect_faces():
    """Detect all faces in image and return embeddings"""
    try:
        if 'image' not in request.files:
            return jsonify({'success': False, 'error': 'No image provided'}), 400

        file = request.files['image']
        
        img_data = file.read()
        img_array = np.frombuffer(img_data, np.uint8)
        img = cv2.imdecode(img_array, cv2.IMREAD_COLOR)

        if img is None:
            return jsonify({'success': False, 'error': 'Invalid image'}), 400

        faces = ctx.get(img)

        result_faces = []
        for face in faces:
            result_faces.append({
                'bbox': face['bbox'].tolist(),
                'confidence': float(face['det_score']),
                'embedding': face['embedding'].tolist()
            })

        return jsonify({
            'success': True,
            'faces': result_faces
        })

    except Exception as e:
        return jsonify({'success': False, 'error': str(e)}), 500


@app.route('/health', methods=['GET'])
def health():
    """Health check endpoint"""
    return jsonify({'status': 'healthy'}), 200


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=False)
