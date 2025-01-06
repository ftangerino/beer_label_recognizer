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

from flask import Flask
from app.routes.routes import ocr_bp

####################################################################################################
# 🔶 MAIN FUNCTION
####################################################################################################

# 🟢 [GENERAL] INITIALIZE FLASK APPLICATION
app = Flask(__name__)

# 🟢 [GENERAL] REGISTER OCR BLUEPRINT ROUTES
app.register_blueprint(ocr_bp)

if __name__ == '__main__':
    # 🟢 [GENERAL] RUN FLASK APP ON PORT 5000, ACCESSIBLE FROM ALL NETWORK INTERFACES
    app.run(host='0.0.0.0', port=5000)
