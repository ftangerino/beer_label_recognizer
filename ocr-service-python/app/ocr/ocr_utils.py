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

import easyocr
from PIL import Image
import io
import difflib

###################################################################################################
# 🛄 RESOURCES
###################################################################################################

reader = easyocr.Reader(['pt', 'en'], gpu=False)
BRAZILIAN_BEERS = [
    "Skol", "Brahma", "Antarctica", "Bohemia", "Itaipava",
    "Kaiser", "Schin", "Devassa", "Polar", "Heineken"
]

###################################################################################################
# 🔰 AUXILIARY FUNCTIONS
###################################################################################################

def extract_text(image_data):
    """
    Realiza leitura da imagem usando easyocr (mantido como estava).
    """
    try:
        image = Image.open(io.BytesIO(image_data))
        result = reader.readtext(image, detail=0)
        raw_text = " ".join(result)
        return raw_text.strip()
    except Exception as e:
        print(f"Erro ao processar a imagem: {e}")
        return None

def match_beer_brand(text):
    """
    Aplica similaridade para retornar a marca de cerveja
    com >= 65% de semelhança, ou None se não encontrado.
    """
    if not text:
        return None
    best_match = None
    highest_ratio = 0.0
    for brand in BRAZILIAN_BEERS:
        seq = difflib.SequenceMatcher(None, text.lower(), brand.lower())
        ratio = seq.ratio() * 100
        if ratio > highest_ratio:
            highest_ratio = ratio
            best_match = brand
    return best_match if highest_ratio >= 65 else None