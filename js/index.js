const crypto = require('crypto');

function constantTimeCompare(a, b) {
  if (typeof a !== 'string' || typeof b !== 'string') return false;
  const bufA = Buffer.from(a);
  const bufB = Buffer.from(b);
  if (bufA.length !== bufB.length) return false;
  return crypto.timingSafeEqual(bufA, bufB);
}

function escapeTemplateString(str) {
  if (typeof str !== 'string') return '';
  return str
    .replace(/\\/g, '\\\\')
    .replace(/`/g, '\\`')
    .replace(/\${/g, '\\${');
}

function encryptAES256GCM(key, text) {
  if (Buffer.from(key).length !== 32) throw new Error('Key must be 32 bytes');
  const iv = crypto.randomBytes(12);
  const cipher = crypto.createCipheriv('aes-256-gcm', key, iv);
  const encrypted = Buffer.concat([cipher.update(text, 'utf8'), cipher.final()]);
  const tag = cipher.getAuthTag();
  return Buffer.concat([iv, tag, encrypted]).toString('base64');
}

function decryptAES256GCM(key, dataB64) {
  if (Buffer.from(key).length !== 32) throw new Error('Key must be 32 bytes');
  const buf = Buffer.from(dataB64, 'base64');
  if (buf.length < 28) throw new Error('Ciphertext too short');
  const iv = buf.subarray(0, 12);
  const tag = buf.subarray(12, 28);
  const encrypted = buf.subarray(28);
  const decipher = crypto.createDecipheriv('aes-256-gcm', key, iv);
  decipher.setAuthTag(tag);
  return decipher.update(encrypted, undefined, 'utf8') + decipher.final('utf8');
}

module.exports = {
  constantTimeCompare,
  escapeTemplateString,
  encryptAES256GCM,
  decryptAES256GCM
};
