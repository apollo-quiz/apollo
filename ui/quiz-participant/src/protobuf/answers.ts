/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { MultiChoiceQuestion, OpenQuestion, QuestionType, questionTypeFromJSON, questionTypeToJSON } from "./questions";

export const protobufPackage = "common.answers";

export interface MultiChoiceAnswer {
  answer: number;
}

export interface OpenAnswer {
  answer: string;
}

export interface Answer {
  type: QuestionType;
  multiChoiceQuestion?: MultiChoiceQuestion | undefined;
  openQuestion?: OpenQuestion | undefined;
  multiChoiceAnswer?: MultiChoiceAnswer | undefined;
  openAnswer?: OpenAnswer | undefined;
}

function createBaseMultiChoiceAnswer(): MultiChoiceAnswer {
  return { answer: 0 };
}

export const MultiChoiceAnswer = {
  encode(message: MultiChoiceAnswer, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.answer !== 0) {
      writer.uint32(8).int32(message.answer);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MultiChoiceAnswer {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMultiChoiceAnswer();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.answer = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MultiChoiceAnswer {
    return { answer: isSet(object.answer) ? globalThis.Number(object.answer) : 0 };
  },

  toJSON(message: MultiChoiceAnswer): unknown {
    const obj: any = {};
    if (message.answer !== 0) {
      obj.answer = Math.round(message.answer);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MultiChoiceAnswer>, I>>(base?: I): MultiChoiceAnswer {
    return MultiChoiceAnswer.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MultiChoiceAnswer>, I>>(object: I): MultiChoiceAnswer {
    const message = createBaseMultiChoiceAnswer();
    message.answer = object.answer ?? 0;
    return message;
  },
};

function createBaseOpenAnswer(): OpenAnswer {
  return { answer: "" };
}

export const OpenAnswer = {
  encode(message: OpenAnswer, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.answer !== "") {
      writer.uint32(10).string(message.answer);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OpenAnswer {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenAnswer();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.answer = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): OpenAnswer {
    return { answer: isSet(object.answer) ? globalThis.String(object.answer) : "" };
  },

  toJSON(message: OpenAnswer): unknown {
    const obj: any = {};
    if (message.answer !== "") {
      obj.answer = message.answer;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<OpenAnswer>, I>>(base?: I): OpenAnswer {
    return OpenAnswer.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<OpenAnswer>, I>>(object: I): OpenAnswer {
    const message = createBaseOpenAnswer();
    message.answer = object.answer ?? "";
    return message;
  },
};

function createBaseAnswer(): Answer {
  return {
    type: 0,
    multiChoiceQuestion: undefined,
    openQuestion: undefined,
    multiChoiceAnswer: undefined,
    openAnswer: undefined,
  };
}

export const Answer = {
  encode(message: Answer, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.type !== 0) {
      writer.uint32(8).int32(message.type);
    }
    if (message.multiChoiceQuestion !== undefined) {
      MultiChoiceQuestion.encode(message.multiChoiceQuestion, writer.uint32(18).fork()).ldelim();
    }
    if (message.openQuestion !== undefined) {
      OpenQuestion.encode(message.openQuestion, writer.uint32(26).fork()).ldelim();
    }
    if (message.multiChoiceAnswer !== undefined) {
      MultiChoiceAnswer.encode(message.multiChoiceAnswer, writer.uint32(34).fork()).ldelim();
    }
    if (message.openAnswer !== undefined) {
      OpenAnswer.encode(message.openAnswer, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Answer {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAnswer();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.type = reader.int32() as any;
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.multiChoiceQuestion = MultiChoiceQuestion.decode(reader, reader.uint32());
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.openQuestion = OpenQuestion.decode(reader, reader.uint32());
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.multiChoiceAnswer = MultiChoiceAnswer.decode(reader, reader.uint32());
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.openAnswer = OpenAnswer.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Answer {
    return {
      type: isSet(object.type) ? questionTypeFromJSON(object.type) : 0,
      multiChoiceQuestion: isSet(object.multiChoiceQuestion)
        ? MultiChoiceQuestion.fromJSON(object.multiChoiceQuestion)
        : undefined,
      openQuestion: isSet(object.openQuestion) ? OpenQuestion.fromJSON(object.openQuestion) : undefined,
      multiChoiceAnswer: isSet(object.multiChoiceAnswer)
        ? MultiChoiceAnswer.fromJSON(object.multiChoiceAnswer)
        : undefined,
      openAnswer: isSet(object.openAnswer) ? OpenAnswer.fromJSON(object.openAnswer) : undefined,
    };
  },

  toJSON(message: Answer): unknown {
    const obj: any = {};
    if (message.type !== 0) {
      obj.type = questionTypeToJSON(message.type);
    }
    if (message.multiChoiceQuestion !== undefined) {
      obj.multiChoiceQuestion = MultiChoiceQuestion.toJSON(message.multiChoiceQuestion);
    }
    if (message.openQuestion !== undefined) {
      obj.openQuestion = OpenQuestion.toJSON(message.openQuestion);
    }
    if (message.multiChoiceAnswer !== undefined) {
      obj.multiChoiceAnswer = MultiChoiceAnswer.toJSON(message.multiChoiceAnswer);
    }
    if (message.openAnswer !== undefined) {
      obj.openAnswer = OpenAnswer.toJSON(message.openAnswer);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Answer>, I>>(base?: I): Answer {
    return Answer.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Answer>, I>>(object: I): Answer {
    const message = createBaseAnswer();
    message.type = object.type ?? 0;
    message.multiChoiceQuestion = (object.multiChoiceQuestion !== undefined && object.multiChoiceQuestion !== null)
      ? MultiChoiceQuestion.fromPartial(object.multiChoiceQuestion)
      : undefined;
    message.openQuestion = (object.openQuestion !== undefined && object.openQuestion !== null)
      ? OpenQuestion.fromPartial(object.openQuestion)
      : undefined;
    message.multiChoiceAnswer = (object.multiChoiceAnswer !== undefined && object.multiChoiceAnswer !== null)
      ? MultiChoiceAnswer.fromPartial(object.multiChoiceAnswer)
      : undefined;
    message.openAnswer = (object.openAnswer !== undefined && object.openAnswer !== null)
      ? OpenAnswer.fromPartial(object.openAnswer)
      : undefined;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
