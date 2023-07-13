import fs from 'fs';
import Long from 'long';
import { trivial } from './gen/proto/trivial';
import { trivial as trivial_pbjs } from './gen/proto/trivial.pbjs';
import { types } from './gen/proto/types';
import { types as types_pbjs } from './gen/proto/types.pbjs';
import * as protobufjs from 'protobufjs';

// protobufjs.util.Long = Long;
// protobufjs.configure();

console.log('Encoding test messages');

function writeBuffer(name: string, buffer: Uint8Array) {
  fs.writeFileSync(name, buffer);
  console.log('Wrote', name);
}

function checkEqual(message: object, decoded: object, decodedPbjs: object) {
  const messageJSON = JSON.stringify(message, null, 2);
  const decodedJSON = JSON.stringify(decoded, null, 2);
  const decodedPbjsJSON = JSON.stringify(decodedPbjs, null, 2);
  if (decodedJSON !== decodedPbjsJSON) {
    console.log('Decoded message does not match original message!');
    console.log('Original message (JS):', message);
    console.log('Original message (JSON):', messageJSON);
    console.log('Decoded (protoc-gen-protobufjs):', decodedJSON);
    console.log('Decoded (pbjs):', decodedPbjsJSON);
  }
}

function encode<T extends object>(
  name: string,
  type: any,
  pbjsType: any,
  message: T
) {
  writeBuffer(
    `./out/protoc-gen-protobufjs/${name}.bin`,
    type.encode(message).finish()
  );
  writeBuffer(`./out/pbjs/${name}.bin`, pbjsType.encode(message).finish());

  // Test encode/decode
  const buffer = type.encode(message).finish();
  const decoded = type.decode(buffer);
  const decodedPbjs = pbjsType.decode(buffer);
  checkEqual(message, decoded, decodedPbjs);

  // Test toJSON
  const json = JSON.stringify(decoded.toJSON(), null, 2);
  const jsonPbjs = JSON.stringify(decodedPbjs.toJSON(), null, 2);
  if (json !== jsonPbjs) {
    console.log('toJSON() does not match protobufjs!');
    console.log('protobufjs-cli: ' + jsonPbjs);
    console.log('protoc-gen-protobufjs: ' + json);
  }
}

encode<trivial.ITrivialMessage>(
  'trivial',
  trivial.TrivialMessage,
  trivial_pbjs.TrivialMessage,
  {
    trivialField: 'Hello',
  }
);

encode<types.IAllSimpleTypes>(
  'types',
  types.AllSimpleTypes,
  types_pbjs.AllSimpleTypes,
  {
    double: 1.5,
    float: 2.5,
    int64: new Long(3, 0),
    uint64: new Long(4, 0, true),
    int32: 5,
    fixed64: new Long(6, 0, true),
    fixed32: 7,
    bool: true,
    string: 'Hello world',
    someMessage: types.AllSimpleTypes.SomeMessage.create({
      child: types.AllSimpleTypes.SomeMessage.SomeMessageChild.create({}),
    }),
    bytes: new Uint8Array([12, 0, 1, 2]),
    mapWithInt64: { foo: new Long(3, 5) },

    // TODO: this fails due to fields being out of order, but is otherwise handled OK.
    // Need to debug this and uncomment.

    // repeatedSomeEnumValue: [1, 2],
  }
);
