const path = require('path');
const fs = require('fs');

const files = [
  'DFSOnBST',
  'LRU',
  'SinglyLinkedList',
  'ArrayList',
  'BTBFS',
  'CompareBinaryTrees',
  'DFSOnBST',
  'DFSGraphList',
  'Trie',
  'BFSGraphMatrix',
  'Map',
  'MinHeap',
];

const done = [
  'BTPreOrder',
  'BTInOrder',
  'BTPostOrder',
  'BinarySearchList',
  'BubbleSort',
  'DoublyLinkedList',
  'LinearSearchList',
  'MazeSolver',
  'Queue',
  'QuickSort',
  'Stack',
  'TwoCrystalBalls',
];

const getRandom = () => {
  let value = files[Math.floor(Math.random() * files.length)];
  if (done.indexOf(value) !== -1) value = getRandom();
  return value;
};

const createFile = (fileName) => {
  const filePath = path.join(__dirname, `${fileName}.go`);
  if (!fs.existsSync(filePath)) fs.writeFileSync(filePath, '');

  const testFilePath = path.join(__dirname, `${fileName}_test.go`);
  if (!fs.existsSync(testFilePath)) fs.writeFileSync(testFilePath, '');
};

const createAllFiles = () => files.forEach((file) => createFile(file));

const createRandomFile = () => createFile(getRandom());

////////////////////////////// TO RUN /////////////////////////////////
createRandomFile();
