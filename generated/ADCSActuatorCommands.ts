/**
* ADCSActuatorCommands
* Commands for the attitude control actuators
*/
export interface ADCSActuatorCommands {
/** Commanded reaction wheel speeds (rpm) */
reactionWheelSpeeds: number[];
/** Commanded magnetorquer dipole moments (mA·m²) */
magnetorquerCommands: number[];
/** Timestamp of the actuator commands (ms) */
commandTimestamp: number;
/** Current ADCS control mode (enum) */
controlMode: number;
}

/**
* Creates a default ADCSActuatorCommands object
* @returns A new ADCSActuatorCommands with default values
*/
export function createADCSActuatorCommands(): ADCSActuatorCommands {
return {
reactionWheelSpeeds: Array(4).fill(0),
magnetorquerCommands: Array(3).fill(0),
commandTimestamp: 0,
controlMode: 0
};
}

/**
* Serializes a ADCSActuatorCommands object to an ArrayBuffer
* @param data The ADCSActuatorCommands object to serialize
* @returns An ArrayBuffer containing the serialized data
*/
export function serializeADCSActuatorCommands(data: ADCSActuatorCommands): ArrayBuffer {
const buffer = new ArrayBuffer(19);
const view = new DataView(buffer);
let offset = 0;
for (let i = 0; i < 4; i++) { view.setInt16(offset,
  data.reactionWheelSpeeds[i], true); offset +=2; }
for (let i = 0; i < 3; i++) { view.setInt16(offset,
  data.magnetorquerCommands[i], true); offset +=2; }
  view.setUint32(offset, data.commandTimestamp, true); offset +=4;
  view.setUint8(offset, data.controlMode); offset +=1; return buffer; } /** * Deserializes an ArrayBuffer to a ADCSActuatorCommands object * @param buffer The ArrayBuffer
  containing serialized data * @returns A ADCSActuatorCommands object with the deserialized data */ export function
  deserializeADCSActuatorCommands(buffer: ArrayBuffer): ADCSActuatorCommands { const view=new DataView(buffer); let offset=0; const
  result=createADCSActuatorCommands(); const reactionWheelSpeedsArray=[]; for (let i=0; i < 4;
  i++) { reactionWheelSpeedsArray.push(view.getInt16(offset, true)); offset +=2; } result.reactionWheelSpeeds=reactionWheelSpeedsArray; const magnetorquerCommandsArray=[]; for (let i=0; i < 3;
  i++) { magnetorquerCommandsArray.push(view.getInt16(offset, true)); offset +=2; } result.magnetorquerCommands=magnetorquerCommandsArray; result.commandTimestamp=view.getUint32(offset, true); offset +=4;
  result.controlMode=view.getUint8(offset); offset +=1; return result; } 