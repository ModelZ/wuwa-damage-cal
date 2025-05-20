# WuWa Damage Calculator

โปรเจกต์นี้เป็นเครื่องคำนวณดาเมจเบื้องต้นสำหรับเกม โดยอ้างอิงสูตรจาก [dmg_formula.txt](dmg_formula.txt)
For English [here](readme-EN.md)

## โครงสร้างโปรเจกต์

- `main.go`: จุดเริ่มต้นของโปรแกรม
- `mz_wuwa_cal/mz_wuwa_cal.go`: โครงสร้างและฟังก์ชันหลักสำหรับคำนวณดาเมจ
- `dmg_formula.txt`: สูตรและคำอธิบายการคำนวณดาเมจ
- `input.toml`: ตัวอย่างไฟล์อินพุตสำหรับกำหนดค่าสถานะผู้โจมตีและเป้าหมาย
- `readme.md`: เอกสารประกอบโปรเจกต์

## สิ่งที่ต้องเตรียม
- download ภาษาโปรแกรมที่ใช้ [Go](https://go.dev/doc/install)

## วิธีใช้งาน

1. **แก้ไขไฟล์ `input.toml`**  
   กำหนดค่าสถานะของผู้โจมตีและเป้าหมายในไฟล์ `input.toml`  
   ตัวอย่าง:
   ```toml
   [attacker]
   # ชื่อผู้โจมตี
   Name = "RoverAero"
   # เลเวล
   Level = 90
   # % ดาเมจสกิล (35.31%)
   AbilityDmg = 0.3531
   # พลังโจมตีรวม (ที่แสดงในหน้าสถานะตัวละคร)
   TotalAtk = 1183
   # ให้ติดคริไหม false=ติดคริ true=ไม่ติดคริ
   Crit = false
   # % โบนัสความเสียหายธาตุ (Spectro/Aero/Fusion/Havoc/Glacio)
   ElementDmgBonus = 0.0
   # % โบนัสความเสียหายประเภทการโจมตี (basic/heavy/reson_skill/reson_riberation)
   AttackTypeDmgBonus = 0.0
   # % final damage (คล้ายๆ outro verina)
   DmgAmplify = 0.0
   # % คริดาเมจ (150%)
   CritDmg = 1.5
   # ไม่สนพลังป้องกัน %
   DefIgnore = 0.0
   # ไม่สนต้านทานธาตุ %
   ElementResIgnore = 0.0
   # โบนัสดาเมจพิเศษ
   SpecialDmg = 0.0

   [target]
   # ชื่อมอนสเตอร์
   Name = "Mourning Aix"
   # เลเวลมอนสเตอร์
   Level = 85
   # ความต้านทานธาตุ %
   BaseElementRes = 0.1
   # ถูกสร้างความเสียหายเพิ่มเติม %
   DmgTaken = 0.0
   # ความต้านทานความเสียหายพื้นฐาน
   DmgResBase = 0.0
   # ความต้านทานความเสียหายเพิ่มเติม
   DmgResAdditional = 0.0
   # ความต้านทานธาตุพื้นฐาน
   ElementResBase = 0.0
   # ความต้านทานธาตุเพิ่มเติม
   ElementResAdditional = 0.0
   # โบนัสดาเมจที่ได้รับ
   DmgAmplify = 0.0
   ```

2. **สร้างโปรแกรม:**
   ```sh
   go mod tidy
   go build main.go
   ```

## คุณสมบัติ

- คำนวณดาเมจจริงจากค่าสถานะของผู้โจมตีและเป้าหมาย
- โครงสร้างโปรแกรมแบบแยกส่วน ขยายต่อได้ง่าย
- อินพุตแบบ TOML ที่อ่านง่ายและใส่คอมเมนต์ได้

## อ้างอิง

ดูรายละเอียดสูตรการคำนวณได้ที่ [dmg_formula.txt](dmg_formula.txt)