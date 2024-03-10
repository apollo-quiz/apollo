/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "common.questions";

export enum QuestionType {
  NONE = 0,
  MULTI_CHOICE = 2,
  OPEN_ANSWER = 1,
  UNRECOGNIZED = -1,
}

export function questionTypeFromJSON(object: any): QuestionType {
  switch (object) {
    case 0:
    case "NONE":
      return QuestionType.NONE;
    case 2:
    case "MULTI_CHOICE":
      return QuestionType.MULTI_CHOICE;
    case 1:
    case "OPEN_ANSWER":
      return QuestionType.OPEN_ANSWER;
    case -1:
    case "UNRECOGNIZED":
    default:
      return QuestionType.UNRECOGNIZED;
  }
}

export function questionTypeToJSON(object: QuestionType): string {
  switch (object) {
    case QuestionType.NONE:
      return "NONE";
    case QuestionType.MULTI_CHOICE:
      return "MULTI_CHOICE";
    case QuestionType.OPEN_ANSWER:
      return "OPEN_ANSWER";
    case QuestionType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface MultiChoiceQuestion {
  points: number;
  option: string[];
  description: string;
}

export interface OpenQuestion {
  points: number;
  description: string;
}

export interface Question {
  type: QuestionType;
  multiChoiceQuestion?: MultiChoiceQuestion | undefined;
  openQuestion?: OpenQuestion | undefined;
}

function createBaseMultiChoiceQuestion(): MultiChoiceQuestion {
  return { points: 0, option: [], description: "" };
}

export const MultiChoiceQuestion = {
  encode(message: MultiChoiceQuestion, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.points !== 0) {
      writer.uint32(8).int32(message.points);
    }
    for (const v of message.option) {
      writer.uint32(18).string(v!);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MultiChoiceQuestion {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMultiChoiceQuestion();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.points = reader.int32();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.option.push(reader.string());
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.description = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MultiChoiceQuestion {
    return {
      points: isSet(object.points) ? globalThis.Number(object.points) : 0,
      option: globalThis.Array.isArray(object?.option) ? object.option.map((e: any) => globalThis.String(e)) : [],
      description: isSet(object.description) ? globalThis.String(object.description) : "",
    };
  },

  toJSON(message: MultiChoiceQuestion): unknown {
    const obj: any = {};
    if (message.points !== 0) {
      obj.points = Math.round(message.points);
    }
    if (message.option?.length) {
      obj.option = message.option;
    }
    if (message.description !== "") {
      obj.description = message.description;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MultiChoiceQuestion>, I>>(base?: I): MultiChoiceQuestion {
    return MultiChoiceQuestion.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MultiChoiceQuestion>, I>>(object: I): MultiChoiceQuestion {
    const message = createBaseMultiChoiceQuestion();
    message.points = object.points ?? 0;
    message.option = object.option?.map((e) => e) || [];
    message.description = object.description ?? "";
    return message;
  },
};

function createBaseOpenQuestion(): OpenQuestion {
  return { points: 0, description: "" };
}

export const OpenQuestion = {
  encode(message: OpenQuestion, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.points !== 0) {
      writer.uint32(8).int32(message.points);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OpenQuestion {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenQuestion();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.points = reader.int32();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.description = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): OpenQuestion {
    return {
      points: isSet(object.points) ? globalThis.Number(object.points) : 0,
      description: isSet(object.description) ? globalThis.String(object.description) : "",
    };
  },

  toJSON(message: OpenQuestion): unknown {
    const obj: any = {};
    if (message.points !== 0) {
      obj.points = Math.round(message.points);
    }
    if (message.description !== "") {
      obj.description = message.description;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<OpenQuestion>, I>>(base?: I): OpenQuestion {
    return OpenQuestion.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<OpenQuestion>, I>>(object: I): OpenQuestion {
    const message = createBaseOpenQuestion();
    message.points = object.points ?? 0;
    message.description = object.description ?? "";
    return message;
  },
};

function createBaseQuestion(): Question {
  return { type: 0, multiChoiceQuestion: undefined, openQuestion: undefined };
}

export const Question = {
  encode(message: Question, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.type !== 0) {
      writer.uint32(8).int32(message.type);
    }
    if (message.multiChoiceQuestion !== undefined) {
      MultiChoiceQuestion.encode(message.multiChoiceQuestion, writer.uint32(18).fork()).ldelim();
    }
    if (message.openQuestion !== undefined) {
      OpenQuestion.encode(message.openQuestion, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Question {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuestion();
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
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Question {
    return {
      type: isSet(object.type) ? questionTypeFromJSON(object.type) : 0,
      multiChoiceQuestion: isSet(object.multiChoiceQuestion)
        ? MultiChoiceQuestion.fromJSON(object.multiChoiceQuestion)
        : undefined,
      openQuestion: isSet(object.openQuestion) ? OpenQuestion.fromJSON(object.openQuestion) : undefined,
    };
  },

  toJSON(message: Question): unknown {
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
    return obj;
  },

  create<I extends Exact<DeepPartial<Question>, I>>(base?: I): Question {
    return Question.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Question>, I>>(object: I): Question {
    const message = createBaseQuestion();
    message.type = object.type ?? 0;
    message.multiChoiceQuestion = (object.multiChoiceQuestion !== undefined && object.multiChoiceQuestion !== null)
      ? MultiChoiceQuestion.fromPartial(object.multiChoiceQuestion)
      : undefined;
    message.openQuestion = (object.openQuestion !== undefined && object.openQuestion !== null)
      ? OpenQuestion.fromPartial(object.openQuestion)
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
