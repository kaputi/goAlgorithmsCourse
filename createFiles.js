const files = [
  'DFSOnBST',
  'LRU',
  'LinearSearchList',
  'BinarySearchList',
  'TwoCrystalBalls',
  'BubbleSort',
  'SinglyLinkedList',
  'DoublyLinkedList',
  'Queue',
  'Stack',
  'ArrayList',
  'MazeSolver',
  'QuickSort',
  'BTPreOrder',
  'BTInOrder',
  'BTPostOrder',
  'BTBFS',
  'CompareBinaryTrees',
  'DFSOnBST',
  'DFSGraphList',
  'Trie',
  'BFSGraphMatrix',
  'Map',
  'MinHeap',
];

const path = require('path');
const fs = require('fs');

files.forEach((file) => {
  // const dirName = path.join(__dirname, file);
  // if (!fs.existsSync(dirName)) fs.mkdirSync(dirName);

  // const fileName = path.join(dirName, `${file}.go`);
  // if (!fs.existsSync(fileName)) fs.writeFileSync(fileName, '');
  // const testFileName = path.join(dirName, `${file}_test.go`);
  // if (!fs.existsSync(testFileName)) fs.writeFileSync(testFileName, '');

  const fileName = path.join(__dirname, `${file}.go`);
  if (!fs.existsSync(fileName)) fs.writeFileSync(fileName, '');

  const testFileName = path.join(__dirname, `${file}_test.go`);
  if (!fs.existsSync(testFileName)) fs.writeFileSync(testFileName, '');
});
