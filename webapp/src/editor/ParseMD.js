import katex from 'katex'

// Regex list
var RegExp = {
  header : /^(#{1,6} )(.*)$/,
  line: /^([-*]{3,})$/,
  beginCode: /^(```)(.*)$/,
  endCode: /^(```)$/,
  tabCode: /^(?:[\t]|[ ]{4})(.*)$/,
  markListBegin: /^([*\-+] )(.*)$/,
  markList: /^([ ]*)([*\-+] )(.*)$/,
  numListBegin: /^([\d]+\. )(.*)$/,
  numList: /^([\t ]*)([\d]+\. )(.*)$/,
  quote: /^(> )(.*)$/,
  linksList: /^\[(.*)\]: (.*)$/,
  image: /^!\[(.*)\]\((.*)\)$/,
  imageList: /^!\[(.*)\]\[(.*)\]$/,
  imageLink: /^\[!\[(.*)\]\((.*)\)\]\((.*)\)$/,
  table: /^\|(.*)\|$/,
  tableOpt: /^\|(?: *:?-{3,}:? *\|)+$/,
  //formula: /^/,
}

// Parsing markdown inside strings
function internalParser(str, linksMap) {
  return str
      .replace(/\$\$([^$]+)\$\$/, function(match, p1) {
        return '<br><div class="text-center">' + katex.renderToString(p1, {
          throwOnError: false,
          output: 'html'
        }) + '</div><br>'
      })
      .replace(/\\\\\((.+?)\\\\\)/g, function(match, p1) {
        console.log(p1)
        return katex.renderToString(p1, {
          throwOnError: false,
          output: 'html'
        })
      })
      .replace(/``([^`]+)``/g, '<kbd>$1</kbd>')
      .replace(/\*\*\*([^*]+)\*\*\*/g, '<i><b>$1</b><i>')
      .replace(/___([^_]+)___/g, '<i><b>$1</b></i>')
      .replace(/\*\*([^*]+)\*\*/g, '<b>$1</b>')
      .replace(/__([^_]+)__/g, '<b>$1</b>')
      .replace(/\*([^*]+)\*/g, '<i>$1</i>')
      .replace(/_([^_]+)_/g, '<i>$1</i>')
      .replace(/~~([^~]+)~~/g, '<s>$1</s>')
      .replace(/\[([^\]]+)]\(([^(]+?)\)/g, '<a href=$2>$1</a>')
      .replace(/\[([^\]]+)]\[([^\]]+)]/g, function(match, p1, p2) {
        if(linksMap.get(p2) != null) {
          return '<a href='+linksMap.get(p2)+'>'+p1+'</a>';
        } else {
          return '<a href=#>'+p1+'</a>';
        }
      })
}

function alignTable(str) {
  var res = str.match(/^-{3,}$/)
  if(res != null) { return 'text-left'}
  res = str.match(/^:-{3,}$/)
  if(res != null) { return 'text-left'}
  res = str.match(/^-{3,}:$/)
  if(res != null) { return 'text-right'}
  res = str.match(/^:-{3,}:$/)
  if(res != null) { return 'text-center'}
}

export default {
  ParseMD: function (text) {

    var arrayComponents = [],
      arrTemp = [],
      i = 0,
      j = 0,
      k = 0,
      links = new Map(),
      result,
      arr = text
        .replace(/^\s+|\r|\s+$/g, '')
        .replace(/\t/g, '    ')
        .split(/\n/)

    // Аssembly of links and blank lines
    for(i = 0; i < arr.length; i++) {
      result = arr[i].match(RegExp.linksList)
      if(result != null) {
        links.set(result[1], result[2])
      } else if (arr[i] == "" && k == 0){
        arrTemp.push(arr[i])
        k++
      } else if(arr[i] != ""){
        arrTemp.push(arr[i])
        k = 0
      }
    }
    
    arr = arrTemp

    for(i = 0; i < arr.length; i++) {

      // Checking tables and tab code
      if(arr[i] == "" && i + 1 < arr.length) {
        j = i + 1

        // TAB CODE
        result = arr[j].match(RegExp.tabCode)
        if(result != null) {
          var list = []
          list.push(result[1])
          for(j+=1; j < arr.length; j++) {
            result = arr[j].match(RegExp.tabCode)
            if(result == null){
              break
            }
            list.push(result[1])
          }
          if(j <= arr.length && j != i + 1) {
            arrayComponents.push({
              name: 'code',
              list: list,
            })
          }
          i = j - 1
          continue
        }

        // TABLE
        result = arr[j].match(RegExp.table)
        if(result != null && j + 1 < arr.length) {
          var headTable = result[1].split('|')
          j++
          result = arr[j].match(RegExp.tableOpt)
          if(result == null) {
            continue
          }

          var arrOpt = result[0].replace(/^\|/, '').replace(/\|$/, '').replace(/ /g, '').split('|')
          if(arrOpt.length != headTable.length) {continue}
          
          var table = {header: [], items: []}
          for(k = 0; k < headTable.length; k++) {
            table.header.push({
              text: headTable[k],
              align: alignTable(arrOpt[k]),
            })
          }
          j++
          while(j < arr.length) {
            result = arr[j].match(RegExp.table)
            if(result == null) {
              break
            }
            var items = result[1].split('|')
            if(items.length != headTable.length) {
              break
            }
            var items1 = []
            for(var z = 0; z < items.length; z++) {
              items1.push(internalParser(items[z], links))
            }
            table.items.push(items1)
            j++
          }
          if(j <= arr.length) {
            arrayComponents.push({
              name: 'table',
              table: table,
            })
          }
          i = j - 1
          continue
        }
        
        // PARAG NEWLINE
        continue
      }

      // HEADER
      result = arr[i].match(RegExp.header)
      if(result != null) {
        arrayComponents.push({
          name: 'header',
          value: internalParser(result[2], links),
          size: result[1].length - 1,
        })
        continue
      }

      // HR
      result = arr[i].match(RegExp.line)
      if(result != null) {
        arrayComponents.push({
          name: 'line'
        })
        continue
      }
      
      // CODE
      result = arr[i].match(RegExp.beginCode)
      if(result != null) {
        list = []
        j=i+1
        while(j < arr.length) {
          result = arr[j].match(RegExp.endCode)
          if(result != null) {
            break
          }
          list.push(arr[j])
          j++
        }
        if(j <= arr.length && j != i + 1) {
          arrayComponents.push({
            name: 'code',
            list: list,
          })
        }
        i = j
        continue
      }

      // MarkLIST
      result = arr[i].match(RegExp.markListBegin)
      if(result != null) {
        list = []
        j = i + 1
        var item = {value: result[2], list:[]}
        k = 0
        while(j < arr.length) {
          result = arr[j].match(RegExp.markList)
          if(result == null) {
            break
          }

          if(result[1].length/4 == 0) {
            list.push(item)
            item = {value: internalParser(result[3], links), list:[]}
            k = 0
          } else if(result[1].length/4 == 1) {
            item.list.push({value: internalParser(result[3], links), list:[]})
            k++
          } else if(result[1].length/4 >= 2 && k > 0) {
            item.list[k - 1].list.push({value: internalParser(result[3], links), list:[]})
          }
          j++
        }
        list.push(item)
        if(j <= arr.length) {
          arrayComponents.push({
            name: 'marklist',
            list: list,
          })
        }
        i = j - 1
        continue
      }

      // NumLIST
      result = arr[i].match(RegExp.numListBegin)
      if(result != null) {
        list = []
        j = i + 1
        item = {value: result[2], list:[]}
        k = 0
        while(j < arr.length) {
          result = arr[j].match(RegExp.numList)
          if(result == null) {
            break
          }
          if(result[1].length/4 == 0) {
            list.push(item)
            item = {value: internalParser(result[3], links), list:[]}
            k = 0
          } else if(result[1].length/4 == 1) {
            item.list.push({value: internalParser(result[3], links), list:[]})
            k++
          } else if(result[1].length/4 >= 2 && k > 0) {
            item.list[k - 1].list.push({value: internalParser(result[3], links), list:[]})
          }
          j++
        }
        list.push(item)
        if(j <= arr.length) {
          arrayComponents.push({
            name: 'numlist',
            list: list,
          })
        }
        i = j - 1
        continue
      }

      // QUOTE
      result = arr[i].match(RegExp.quote)
      if(result != null) {
        list = []
        j = i + 1
        list.push({value:result[2]}) 
        while(j < arr.length) {
          result = arr[j].match(RegExp.quote)
          if(result == null) {
            break
          }
          list.push({value:internalParser(result[2], links)})
          j++
        }
        if(j <= arr.length) {
          arrayComponents.push({
            name: "quote",
            list: list,
          })
        }
        i=j-1
        continue
      }

      // IMAGE
      result = arr[i].match(RegExp.image)
      if(result != null) {
        arrayComponents.push({
          name: 'image',
          alt: result[1],
          value: result[2],
        })
        continue
      }

      result = arr[i].match(RegExp.imageList)
      if(result != null) {
        if(links.get(result[2]) != null) {
          arrayComponents.push({
            name: 'image',
            alt: result[1],
            value: links.get(result[2]),
          })
        } else {
          arrayComponents.push({
            name: 'image',
            alt: result[1],
            value: '',
          })
        }
        continue
      }

      // Image Link
      result = arr[i].match(RegExp.imageLink)
      if(result != null) {
        arrayComponents.push({
          name: 'imagelink',
          alt: result[1],
          value: result[2],
          link: result[3],
        })
        continue
      }

      
      // TEXT
      arrayComponents.push({
        name: 'paragraph',
        value: internalParser(arr[i], links)
      })

    }

    console.log(arrayComponents)
    return arrayComponents
  },

}