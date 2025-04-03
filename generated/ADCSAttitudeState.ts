/**
* ADCSAttitudeState
* Current attitude state of the spacecraft
*/
export interface ADCSAttitudeState {
/** Quaternion representing the spacecraft attitude (quaternion) */
quaternion: number[];
/** Angular velocity vector of the spacecraft (rad/s) */
angularVelocity: number[];
/** Timestamp of the attitude measurement (ms) */
timestamp: number;
/** Current mode of attitude determination (enum) */
attitudeDeterminationMode: number;
/** Flag indicating if the attitude solution is valid */
attitudeValid: boolean;
}

/**
* Creates a default ADCSAttitudeState object
* @returns A new ADCSAttitudeState with default values
*/
export function createADCSAttitudeState(): ADCSAttitudeState {
return {
quaternion: Array(4).fill(0),
angularVelocity: Array(3).fill(0),
timestamp: 0,
attitudeDeterminationMode: 0,
attitudeValid: false
};
}

/**
* Serializes a ADCSAttitudeState object to an ArrayBuffer
* @param data The ADCSAttitudeState object to serialize
* @returns An ArrayBuffer containing the serialized data
*/
export function serializeADCSAttitudeState(data: ADCSAttitudeState): ArrayBuffer {
const buffer = new ArrayBuffer(34);
const view = new DataView(buffer);
let offset = 0;
for (let i = 0; i < 4; i++) { view.setFloat32(offset, data.quaternion[i], true); offset +=4; }
for (let i = 0; i < 3; i++) { view.setFloat32(offset, data.angularVelocity[i], true); offset +=4; }
  view.setUint32(offset, data.timestamp, true); offset +=4;
  view.setUint8(offset, data.attitudeDeterminationMode); offset +=1; view.setUint8(offset,
  data.attitudeValid ? 1 : 0); offset +=1; return buffer; } /** * Deserializes an ArrayBuffer to a ADCSAttitudeState object * @param buffer The ArrayBuffer
  containing serialized data * @returns A ADCSAttitudeState object with the deserialized data */ export function
  deserializeADCSAttitudeState(buffer: ArrayBuffer): ADCSAttitudeState { const view=new DataView(buffer); let offset=0; const
  result=createADCSAttitudeState(); const quaternionArray=[]; for (let i=0; i < 4;
  i++) {
  quaternionArray.push(view.getFloat32(offset, true)); offset +=4; } result.quaternion=quaternionArray; const angularVelocityArray=[]; for (let i=0; i < 3;
  i++) {
  angularVelocityArray.push(view.getFloat32(offset, true)); offset +=4; } result.angularVelocity=angularVelocityArray; result.timestamp=view.getUint32(offset, true); offset +=4;
  result.attitudeDeterminationMode=view.getUint8(offset); offset +=1; result.attitudeValid=view.getUint8(offset) !==0;
  offset +=1; return result; } 