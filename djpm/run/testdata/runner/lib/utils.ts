export class Utils {
    static alloc(size: i32): i32 {
        return heap.alloc(size) as i32;
    }

    static stringToUint8Arrary(s: string): Uint8Array {
        const buffer = String.UTF8.encode(s);
        if (buffer.byteLength === 0) {
            return new Uint8Array(0);
        }

        return Uint8Array.wrap(buffer, 0, s.length);
    }

    static uint8ArrayToString(arr: Uint8Array): string {
        return String.UTF8.decode(arr.buffer, false);
    }

    static uint8ArrayToHex(data: Uint8Array): string {
        let hexString = "";
        for (let i = 0; i < data.length; i++) {
            let hex = data[i].toString(16);
            if (hex.length < 2) {
                hex = "0" + hex;
            }
            hexString += hex;
        }
        return hexString;
    }

    static hexToUint8Array(s: string): Uint8Array {
        s = s.replace(/\s/g, "");
        if (s.length % 2 !== 0) {
            throw new Error("Invalid hex string");
        }

        const result = new Uint8Array(s.length / 2);

        for (let i = 0, j = 0; i < s.length; i += 2, j++) {
            const byteString = s.substring(i, 2);
            const byte = parseInt(byteString, 16);
            if (isNaN(byte)) {
                throw new Error("Invalid hex string");
            }
            result[j] = byte;
        }

        return result;
    }

    static uint8ArrayToBool(data: Uint8Array): bool {
        for (let i = 0; i < data.length; i++) {
            if (data[i] != 0) {
                return true;
            }
        }
        return false;
    }

    static boolToUint8Array(b: bool): Uint8Array {
        const result = new Uint8Array(1);
        result[0] = b ? 1 : 0;

        return result;
    }

    static concatUint8Arrays(a: Uint8Array, b: Uint8Array): Uint8Array {
        const result = new Uint8Array(a.length + b.length);

        for (let i = 0; i < a.length; i++) {
            result[i] = a[i];
        }

        for (let i = 0; i < b.length; i++) {
            result[a.length + i] = b[i];
        }

        return result;
    }
}