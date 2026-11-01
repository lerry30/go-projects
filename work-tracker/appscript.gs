function doGet() {
  return ContentService.createTextOutput('alive');
}

function doPost(e) {
  try {
    const data = JSON.parse(e.postData.contents);
    
    // verify a shared secret
    // if (data.secret !== 'your-shared-secret-here') {
    //   return ContentService.createTextOutput(JSON.stringify({ error: 'unauthorized' }))
    //     .setMimeType(ContentService.MimeType.JSON);
    // }
    
    // Do whatever you need with the incoming change
    const res = handleDataChange(data);
    
    return ContentService.createTextOutput(JSON.stringify(res))
      .setMimeType(ContentService.MimeType.JSON);
  } catch(err) {
    return ContentService.createTextOutput(JSON.stringify({ error: err.toString() }))
      .setMimeType(ContentService.MimeType.JSON);
  }
}

function handleDataChange(reqBody) {
  const trackerId = reqBody['tracker-id'];
  const sheetName = reqBody['sheet-name'];
  const fname = reqBody['first-name'];
  const action = reqBody['log'];

  const tracker = SpreadsheetApp.openById(trackerId);
  const sheetNames = tracker.getSheets().map(sh => sh.getName());
  const logActions = ['login', '1st-break-out', '1st-break-in', 'lunch-out', 'lunch-in', '2nd-break-out', '2nd-break-in', 'logout'];

  if(!sheetNames.includes(sheetName)) {
    return {status: 'fail: sheet not found', action};
  }

  const sheet = tracker.getSheetByName(sheetName);
  const data = sheet.getDataRange().getValues();

  const colLetters = [];
  let empRowNo;

  data.forEach((row, rowIndex) => {
    let loginRow = false;
    row.forEach((cell, colIndex) => {
      // +1 because getRange() is 1-based, but array index is 0-based
      const range = sheet.getRange(rowIndex + 1, colIndex + 1);
      const notation = range.getA1Notation(); // e.g. "B3"
      const rowNumber = range.getRow();       // e.g. 3
      const colLetter = notation.replace(/[0-9]/g, ""); // e.g. "B"

      //Logger.log(`Cell ${notation} → value: "${cell}", row: ${rowNumber}, col: ${colLetter}`);
      const nCell = String(cell).trim().toLowerCase();

      if(nCell.match(/login/)) {
        loginRow = true;
      }

      if(loginRow &&
      (nCell.match(/login/) || 
      nCell.match(/out/) || 
      nCell.match(/in/) || 
      nCell.match(/logout/))
      ) {
        colLetters.push(colLetter);
      }
      
      if(nCell.match(new RegExp(fname, 'ig'))) {
        //Logger.log(`Cell ${notation} → value: "${cell}", row: ${rowNumber}, col: ${colLetter}`);
        empRowNo = rowNumber;
      }
    });
  });

  if(logActions.length !== colLetters.length) {
    return {status: 'required columns not found', action};
  }

  const i = logActions.indexOf(action);
  if(i < 0) {
    return {status: 'action not found', action};
  }

  const cell = `${colLetters[i]}${empRowNo}`;
  sheet.getRange(cell).setValue(formattedCurrentTime());

  if(sheet.getRange(cell).getValue().trim() === "") {
    return {status: 'failed to insert log', action};
  }

  return {status: 'successfully inserted', action};
}

function formattedCurrentTime() {
  const now = new Date();
  const formatted = now.toLocaleTimeString("en-US", {
    hour: "numeric",
    minute: "2-digit",
    second: "2-digit",
    hour12: true,
  });
  return formatted;
}

/*
const debugSheet = SpreadsheetApp.openById('1X85mZPnWqzp0XfgZgsIR4MJZsl73m2yUCkcROxkTqrk').getSheetByName('log')
    || SpreadsheetApp.openById('1X85mZPnWqzp0XfgZgsIR4MJZsl73m2yUCkcROxkTqrk').insertSheet('log');
  debugSheet.appendRow(['-->', sheet.getRange(cell).getValue()]);
*/




