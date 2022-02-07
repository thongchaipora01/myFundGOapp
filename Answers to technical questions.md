# myFundGOapp
### Technical questions

 * ใช้เวลาทำแบบทดสอบไปเท่าไร ถ้ามีเวลามากกว่านี้จะทำอะไรเพิ่ม ถ้าใช้เวลาน้อยในการทำโจทย์สามารถใช้โอกาสนี้ในการอธิบายได้ว่าอยากเพิ่มอะไร หรือแก้ไขในส่วนไหน
 
   Ans. ใช้เวลาไปทั้งหมด 15 ชั่วโมง ถ้ามีเวลามากกว่านี้ก็อยากพัฒนา UI ให้ดูน่าใช้กว่านี้ หรือ อาจจะต้องใช้ framework ภาษาอื่น(e.g. React, Angular)มาช่วยทำ UI 
   แล้วนำ GO มาทำเป็น API หลังบ้าน
 * อะไรคือ feature ที่นำเข้ามาใช้ในการพัฒนา application นี้ กรุณาแนบ code snippet มาด้วยว่าใช้อย่างไร ในส่วนไหน
 
   Ans. Application ถูกพัฒนาโดยภาษา GO ซึ่งเป็นภาษาที่เหมาะสำหรับทำหลังบ้าน แต่ก็มี library บางตัวของ GO ที่สามารถนำมาสร้างเป็น GUI ที่แสดงใน windows ได้
   อย่าง git ของ https://github.com/gen2brain/dlgs ผมก็ import git ของเขามาทำเป็น GUI ของ application นี้เพื่อมารับ input จาก User และแสดง 
   กองทุนที่ผลตอบแทนดีที่สุดให้กับ User ตามช่วงเวลาที่ User เลือก
   
   ```go
   // make GUI of time range list for User choose
   rangeSelected, _, err := dlgs.List("Choose your interest range", "Select item from list:", []string{"1D", "1W", "1M", "1Y"})
   if err != nil {
        panic(err)
   }
   ```
   ส่วนตัว feature หลัก หลังบ้านส่วนใหญ่ใช้ built-in library ของ GO และใส่ logic ของผมเข้าไป อย่าง func GetSuggestFund(timeRange string) model.ResFundArr 
   ก็จะนำ input timeRange จาก User ไปคำนวณและเรียงลำดับหากองทุนที่ดีที่สุดในช่วงเวลาที่ User เลือกไว้
   
   ```go
   // query Sorted Fund from User's selected timeRange
      fund := serv.GetSuggestFund(rangeSelected)
   ```
   และก็นำ Json ในรูปแบบของ struct มาเรียงลำดับด้วย func sortFund(fundArr model.FundArr) model.FundArr
   
   ```go
   // sort Fund(slice) by Performance
   sortedFund := sortFund(fundArr)
   ```
   นำ struct ที่ถูกเรียงลำดับแล้ว มาหากองทุนที่ดีที่สุดในช่วงเวลาที่ User เลือกไว้ ด้วย 
   func splitFundAndChooseRange(sortedFund model.FundArr, timeRange string) (chosenFund model.FundArr)
   ```go
   // response BEST Fund
   angeFund := splitFundAndChooseRange(sortedFund, timeRange)
   ```
 * เราจะสามารถติดตาม performance issue บน production ได้อย่างไร เคยมีประสบการณ์ด้านนี้ไหม

   Ans. performance issue อาจเกิดได้จากหลายกรณี อย่างที่ผมเคยเจอก็จะเป็นเรื่อง memcache ที่มาจาก library GO cache 
   ซึ่งมันจะผูกไว้กับตัว app process บนเครื่อง linux production ซึ่งถ้า memcache เยอะขึ้น มันก็จะทำให้ memory ของเครื่อง
   เยอะขึ้นตาม จนเครื่องไม่สามารถจัดหา memory ให้ application ได้ จึงทำให้ app ตาย ซึ่งผมก็ทำ crontab ไว้เพื่อ start 
   app ใหม่ และทำ log ของการ start ด้วย crontab ไว้เพื่อดูใน 1 วัน app restart ก็ครั้ง
 * อยากปรับปรุง FINNOMENA APIs ที่ใช้ในการพัฒนา ในส่วนไหนให้ดียิ่งขึ้น

   Ans. ผมไม่ได้นำใช้ FINNOMENA APIs มาพัฒนา app ครับ แต่จะนำ Json จาก
   https://storage.googleapis.com/finno-ex-re-v2-static-staging/recruitment-test/fund-ranking-1Y.json
   มาทำแทนเพื่อที่จะมาเพิ่ม date ให้ได้ timeRange ที่จะ test (e.g. 1D, 1W, 1M, 1Y)
   แต่ผมก็ได้ดู list กองทุนผ่าน web https://www.finnomena.com/fund/filter?page=1 การ response ของ Fundlist api
   ถือว่าเร็วมากๆ แล้วครับ แต่ถ้าอยากปรับปรุงด้านอื่นๆ ผมก็สามารถช่วยได้ครับ ^^
