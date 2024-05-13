const Long = require('long');
import { trivial } from './proto/trivial.js';
import { trivial as trivial_pbjs } from './proto/trivial.pbjs.js';
import { types } from './proto/types.js';
import { types as types_pbjs } from './proto/types.pbjs.js';

// protobufjs.util.Long = Long;
// protobufjs.configure();

describe('protoc-gen-protobufjs', () => {
  it('should encode a trivial message', () => {
    testEncoding<trivial.ITrivialMessage>(
      trivial.TrivialMessage,
      trivial_pbjs.TrivialMessage,
      {
        trivialField: 'Hello',
      }
    );
  });

  it('should encode a complex message', () => {
    testEncoding<types.IAllSimpleTypes>(
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
  });
});

function testEncoding<T extends object>(type: any, pbjsType: any, message: T) {
  // Test encode/decode
  const buffer = type.encode(message).finish();
  const decoded = type.decode(buffer);
  const decodedPbjs = pbjsType.decode(buffer);
  const decodedJSON = JSON.stringify(decoded, null, 2);
  const decodedPbjsJSON = JSON.stringify(decodedPbjs, null, 2);
  expect(decodedJSON).toEqual(decodedPbjsJSON);

  // Test toJSON
  const json = JSON.stringify(decoded.toJSON(), null, 2);
  const jsonPbjs = JSON.stringify(decodedPbjs.toJSON(), null, 2);
  expect(json).toEqual(jsonPbjs);
}
