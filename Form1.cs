using System.Security.Cryptography;



namespace ValidationTools
{

    public partial class Form1 : Form
    {

        public Form1()
        {
            InitializeComponent();
            this.AllowDrop = true;
        }


        public static string ComputeHash(string filePath, HashAlgorithm hashAlgorithm)
        {
            using (var stream = File.OpenRead(filePath))
            {
                var hash = hashAlgorithm.ComputeHash(stream);
                return BitConverter.ToString(hash).Replace("-", "").ToLowerInvariant();
            }
        }


        private void Form1_Load(object sender, EventArgs e)
        {
            this.Text = "取文件数据摘要工具 V1.0";
        }

        private void Form1_DragDrop(object sender, DragEventArgs e)
        {
            string[] files = (string[])e.Data.GetData(DataFormats.FileDrop);
            if (files.Length > 0)
            {
                string filePath = files[0];
                string fileName = Path.GetFileName(filePath);
                string md5 = ComputeHash(filePath, MD5.Create());
                string sha1 = ComputeHash(filePath, SHA1.Create());
                string sha256 = ComputeHash(filePath, SHA256.Create());
                string sha512 = ComputeHash(filePath, SHA512.Create());
                // 设置内容
                textBox1.Text = fileName;
                textBox2.Text = sha1;
                textBox3.Text = md5;
                textBox4.Text = sha256;
                textBox5.Text = sha512;

                MessageBox.Show("完成", "提示", MessageBoxButtons.OK, MessageBoxIcon.Information);
            }
        }

        private void Form1_DragEnter(object sender, DragEventArgs e)
        {
            if (e.Data.GetDataPresent(DataFormats.FileDrop))
            {
                e.Effect = DragDropEffects.Copy;
            }
        }
    }
}
