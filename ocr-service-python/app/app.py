###################################################################################################
# ✅ → Discussed and realized
# 🟢 → Discussed and not realized (to be done after the meeting)
# 🟡 → Little important and not discussed (unhindered)
# 🔴 → Very important and not discussed (hindered)
# ❌ → Canceled
# ⚪ → Postponed (technical debit)
###################################################################################################

from flask import Flask

####################################################################################################
# 🔶 MAIN FUNCTION
####################################################################################################
app = Flask(__name__)
@app.route('/health', methods=['GET'])
def health_check():
    return "OK", 200

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=5001)
