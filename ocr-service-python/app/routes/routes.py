###################################################################################################
# 📅 IMPORTS | CODING: UTF-8
###################################################################################################
# ✅ → Discussed and realized
# 🟢 → Discussed and not realized (to be done after the meeting)
# 🟡 → Little important and not discussed (unhindered)
# 🔴 → Very important and not discussed (hindered)
# ❌ → Canceled
# ⚪ → Postponed (technical debit)
###################################################################################################

from flask import Blueprint, request, jsonify
from app.ocr.ocr_utils import perform_ocr, find_best_match
import cv2
import numpy as np
import threading

# 🟢 [GENERAL] CREATE FLASK BLUEPRINT FOR OCR SERVICE
ocr_bp = Blueprint('ocr', __name__)

# 🟢 [GENERAL] THREAD LOCK FOR SYNCHRONIZATION
lock = threading.Lock()

###################################################################################################
# 🔰 AUXILIARY FUNCTIONS
###################################################################################################

@ocr_bp.route('/ocr', methods=['POST'])
def process_image():
    with lock:
        try:
            # 🟢 [GENERAL] GET IMAGE FILE FROM REQUEST
            file = request.files['file']
            img_array = np.frombuffer(file.read(), np.uint8)
            image = cv2.imdecode(img_array, cv2.IMREAD_COLOR)

            # 🔴 [ERROR HANDLING] CHECK IF IMAGE WAS PROPERLY DECODED
            if image is None:
                return jsonify({"error": "Erro ao processar imagem"}), 500

            # 🟢 [GENERAL] CONVERT IMAGE TO RGB FORMAT
            image = cv2.cvtColor(image, cv2.COLOR_BGR2RGB)

            # 🟢 [GENERAL] PERFORM OCR ON THE IMAGE
            extracted_text, error = perform_ocr(image)

            # 🔴 [ERROR HANDLING] RETURN ERROR IF OCR FAILS
            if error:
                return jsonify({"error": error}), 500

            # 🟢 [GENERAL] FIND BEST MATCH FROM EXTRACTED TEXT
            match = find_best_match(extracted_text)

            # 🟢 [GENERAL] RETURN MATCH RESULT OR NO MATCH MESSAGE
            if match:
                return jsonify({"match": f"Lata de {match}"})
            else:
                return jsonify({"match": "Nenhuma correspondência encontrada."})

        except Exception as e:
            # 🔴 [ERROR HANDLING] RETURN GENERAL EXCEPTION ERROR
            return jsonify({"error": str(e)}), 500
