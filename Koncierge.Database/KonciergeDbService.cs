using LiteDB;
using System.IO;

namespace Koncierge.Database
{
    public class KonciergeDbService: IKonciergeDbService
    {
        private readonly string _envVar="Koncierge_Key";
        private readonly LiteDatabase _db;
        private string _dbPath;
        private string _key;

        public KonciergeDbService() {

            _db = GetDatabase();


        }





        private LiteDatabase GetDatabase() {


            string dbFolder = Environment.GetFolderPath(Environment.SpecialFolder.ApplicationData);

            _dbPath = Path.Combine(dbFolder, "Koncierge.db");

            _key = Environment.GetEnvironmentVariable(_envVar, EnvironmentVariableTarget.User);

            if (_key is null) {
                _key=Guid.NewGuid().ToString();

                Environment.SetEnvironmentVariable(_envVar, _key,EnvironmentVariableTarget.User);

            }
          

            return  new LiteDatabase($"Filename={_dbPath};Password={_key}");




        }


    }
}
