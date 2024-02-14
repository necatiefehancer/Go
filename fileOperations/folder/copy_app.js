const { MongoClient } = require('mongodb');

// MongoDB bağlantı URL'si
const url = 'mongodb://localhost:27017';

// MongoDB istemcisini oluştur
const client = new MongoClient(url, { useUnifiedTopology: true });

// Veritabanına bağlan
async function connectToMongoDB() {
  try {
    await client.connect();
    console.log('MongoDB\'ye başarıyla bağlandı');
  } catch (error) {
    console.error('MongoDB bağlantısı başarısız:', error);
  }
}

// Yeni bir belge ekleyip sorgu yap
async function addAndQueryDocument() {
  const database = client.db('test');
  const collection = database.collection('users');

  // Belge ekle
  await collection.insertOne({ name: 'John Doe', age: 30 });

  // Belge sorgula
  const result = await collection.findOne({ name: 'John Doe' });
  console.log('Sorgu sonucu:', result);
}

// MongoDB'ye bağlan ve işlemleri gerçekleştir
async function main() {
  await connectToMongoDB();
  await addAndQueryDocument();
}

// Ana işlemi başlat
main().finally(() => client.close());
