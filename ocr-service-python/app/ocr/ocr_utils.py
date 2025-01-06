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

from paddleocr import PaddleOCR
from difflib import SequenceMatcher

###################################################################################################
# 🛄 RESOURCES
###################################################################################################

# 🟢 [GENERAL] INITIALIZE OCR INSTANCE FOR REUSE
ocr = PaddleOCR(lang='pt', use_angle_cls=True)

# 🟢 [GENERAL] LIST OF BRAZILIAN BEER BRANDS FOR MATCHING
BRAZILIAN_BEERS = [
    "Skol", "Brahma", "Antarctica", "Bohemia", "Itaipava",
    "Kaiser", "Schin", "Devassa", "Polar", "Heineken"
]

###################################################################################################
# 🔰 AUXILIARY FUNCTIONS
###################################################################################################

# 🟢 [GENERAL] CALCULATE TEXT SIMILARITY BETWEEN STRINGS
def calculate_similarity(a, b):
    return SequenceMatcher(None, a.lower(), b.lower()).ratio()

# 🟢 [GENERAL] FIND BEST MATCH FOR EXTRACTED TEXT AGAINST BEER LIST
def find_best_match(extracted_text):
    best_match = None
    highest_similarity = 0.0
    threshold = 0.65  # 🟢 [CONFIG] MINIMUM SIMILARITY THRESHOLD

    words = extracted_text.split()

    for word in words:
        for beer in BRAZILIAN_BEERS:
            similarity = calculate_similarity(word, beer)
            if similarity > highest_similarity and similarity >= threshold:
                highest_similarity = similarity
                best_match = beer

    return best_match

# 🟢 [GENERAL] PERFORM OCR AND RETRY IF NECESSARY
def perform_ocr(image, max_retries=3):
    for attempt in range(max_retries):
        try:
            # 🟢 [OCR] PERFORM OCR ON IMAGE
            result = ocr.ocr(image, cls=True)
            extracted_text = ' '.join([detection[1][0] for detection in result[0] if detection[1]])
            print(f"Tentativa {attempt + 1}: Texto Extraído (OCR):", extracted_text)
            if extracted_text:
                return extracted_text, None
            else:
                raise ValueError("Nenhum texto extraído.")
        except Exception as e:
            # 🔴 [ERROR HANDLING] LOG FAILURE AND RETRY
            print(f"Tentativa {attempt + 1} falhou:", str(e))
            if attempt == max_retries - 1:
                return None, str(e)
    return None, "Falha após múltiplas tentativas."
