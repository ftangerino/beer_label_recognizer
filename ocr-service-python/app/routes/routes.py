###################################################################################################
# 📥 IMPORTS | CODING: UTF-8
###################################################################################################
# ✅ → Discussed and realized
# 🟢 → Discussed and not realized (to be done after the meeting)
# 🟡 → Little important and not discussed (unhindered)
# 🔴 → Very important and not discussed (hindered)
# ❌ → Canceled
# ⚪ → Postponed (technical debit)
###################################################################################################

from flask import request, jsonify
from app.app import app
from app.ocr.ocr_utils import extract_text, match_beer_brand

###################################################################################################
# 🔰 AUXILIARY FUNCTIONS
###################################################################################################

@app.route('/process-image', methods=['POST'])
def process_image():
    if 'image' not in request.files:
        return jsonify({"error": "Nenhuma imagem enviada"}), 400
    file = request.files['image']
    image_data = file.read()

    raw_text = extract_text(image_data)
    if not raw_text:
        return jsonify({"error": "Não foi possível reconhecer texto"}), 404
    recognized_brand = match_beer_brand(raw_text)
    if not recognized_brand:
        return jsonify({"error": "Nenhuma marca identificada"}), 404

    return jsonify({"brand": recognized_brand}), 200
